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
		token, err := app.getAuthTokenFromHeader(r)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

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

func (app *application) getAuthTokenFromHeader(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	parts := strings.Split(authorization, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("authorization header is malformed")
	}

	return parts[1], nil
}

func (app *application) getUser(ctx context.Context, userID string) (*models.User, error) {
	if !app.config.redisdbCfg.enabled {
		return app.store.Users.FindOne(ctx, userID)
	}

	user, err := app.cacheStore.Users.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = app.store.Users.FindOne(ctx, userID)
		if err != nil {
			return nil, err
		}

		if err := app.cacheStore.Users.Set(ctx, user); err != nil {
			return nil, err
		}
	}
	return user, nil
}
