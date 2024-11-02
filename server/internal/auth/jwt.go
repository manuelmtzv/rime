package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthenticator struct {
	secret        string
	refreshSecret string
	audience      string
	issuer        string
}

func NewJWTAuthenticator(secret, refreshSecret, aud, issuer string) *JWTAuthenticator {
	return &JWTAuthenticator{
		secret,
		refreshSecret,
		aud,
		issuer,
	}
}

func (a *JWTAuthenticator) GenerateToken(claims jwt.Claims, tokenType string) (string, error) {
	secret := a.secret
	if tokenType == "refresh" {
		secret = a.refreshSecret
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *JWTAuthenticator) ValidateToken(token, tokenType string) (*jwt.Token, error) {
	secret := a.secret
	if tokenType == "refresh" {
		secret = a.refreshSecret
	}

	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(a.issuer),
		jwt.WithIssuer(a.issuer),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
