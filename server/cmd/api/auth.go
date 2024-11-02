package main

import (
	"errors"
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

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
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

	userResponse := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	if err := app.jsonResponse(w, http.StatusCreated, userResponse); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
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

	if err := app.jsonResponse(w, http.StatusOK, map[string]string{"token": accessToken, "refreshToken": refreshToken}); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) composeTokens(userID string) (string, string, error) {
	accessTokenClaims := &jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(app.config.auth.jwt.expires).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.jwt.issuer,
		"aud": app.config.auth.jwt.issuer,
	}

	accessToken, err := app.authenticator.GenerateToken(accessTokenClaims, "access")
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := &jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(app.config.auth.jwt.refreshExpires).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.jwt.issuer,
		"aud": app.config.auth.jwt.issuer,
	}

	refreshToken, err := app.authenticator.GenerateToken(refreshTokenClaims, "refresh")
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
