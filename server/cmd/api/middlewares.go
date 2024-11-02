package main

import (
	"context"
	"fmt"
	"net/http"
	"rime-api/internal/models"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (app *application) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
			return
		}

		parts := strings.Split(authorization, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
			return
		}

		token := parts[1]

		accessToken, err := app.authenticator.ValidateToken(token, "access")
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		claims, _ := accessToken.Claims.(jwt.MapClaims)
		userID := fmt.Sprintf("%s", claims["sub"])
		ctx := r.Context()

		user, err := app.getUser(ctx, userID)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) getUser(ctx context.Context, userID string) (*models.User, error) {
	return app.store.Users.FindOne(ctx, userID)
}
