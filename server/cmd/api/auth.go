package main

import (
	"errors"
	"fmt"
	"net/http"
	"rime-api/internal/hash"
	"rime-api/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterPayload struct {
	Name     string `json:"name" validate:"required,max=100"`
	Lastname string `json:"lastname" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	User         *UserResponse `json:"user"`
	Token        string        `json:"token"`
	RefreshToken string        `json:"refreshToken"`
}

func (app *application) validate(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromContext(r)

	if user == nil {
		app.unauthorizedErrorResponse(w, r, errors.New("your session could not be validated"))
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, map[string]*models.User{"user": user}); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
	token, err := app.getAuthTokenFromHeader(r)
	if err != nil {
		app.unauthorizedErrorResponse(w, r, err)
		return
	}

	refreshToken, err := app.authenticator.ValidateToken(token, "refresh")

	if err != nil {
		app.unauthorizedErrorResponse(w, r, err)
		return
	}

	claims, _ := refreshToken.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%s", claims["sub"])

	user, err := app.getUser(r.Context(), userID)

	if err != nil {
		app.unauthorizedErrorResponse(w, r, err)
		return
	}

	accessToken, err := app.composeToken(user.ID, "access")
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, map[string]string{"token": accessToken, "refreshToken": refreshToken.Raw}); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("error reading JSON in register: %w", err))
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if user, err := app.store.Users.FindByIdentifier(r.Context(), payload.Email); err == nil && user != nil {
		app.badRequestResponse(w, r, errors.New("email already exists"))
		return
	}

	hashedPassword, err := hash.HashPassword(payload.Password)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	user := &models.User{
		Name:     payload.Name,
		Lastname: payload.Lastname,
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
	}

	if err := app.store.Users.Create(r.Context(), user); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	accessToken, refreshToken, err := app.composeTokens(user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	userResponse := &AuthResponse{
		User: &UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Lastname:  user.Lastname,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
		Token:        accessToken,
		RefreshToken: refreshToken,
	}

	if err := app.jsonResponse(w, http.StatusCreated, userResponse); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("error reading JSON in login: %w", err))
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.store.Users.FindByIdentifier(r.Context(), payload.Username)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if user == nil {
		app.notFoundResponse(w, r, errors.New("user not found"))
		return
	}

	if valid, _ := hash.VerifyPassword(payload.Password, user.Password); !valid {
		app.badRequestResponse(w, r, errors.New("invalid password"))
		return
	}

	accessToken, refreshToken, err := app.composeTokens(user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	authResponse := &AuthResponse{
		User: &UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Lastname:  user.Lastname,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
		Token:        accessToken,
		RefreshToken: refreshToken,
	}

	if err := app.jsonResponse(w, http.StatusOK, authResponse); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) composeTokens(userID string) (accessToken string, refreshToken string, err error) {
	tokens := map[string]*string{
		"access":  &accessToken,
		"refresh": &refreshToken,
	}

	for tokenType, tokenRef := range tokens {
		*tokenRef, err = app.composeToken(userID, tokenType)
		if err != nil {
			return "", "", err
		}
	}

	return accessToken, refreshToken, nil
}

func (app *application) composeToken(userID string, tokenType string) (string, error) {
	expires := app.config.auth.jwt.expires
	if tokenType == "refresh" {
		expires = app.config.auth.jwt.refreshExpires
	}

	accessTokenClaims := &jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(expires).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.jwt.issuer,
		"aud": app.config.auth.jwt.issuer,
	}

	accessToken, err := app.authenticator.GenerateToken(accessTokenClaims, tokenType)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
