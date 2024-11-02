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

func (a *JWTAuthenticator) generateToken(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *JWTAuthenticator) GenerateAccessToken(claims jwt.Claims) (string, error) {
	return a.generateToken(claims, a.secret)
}

func (a *JWTAuthenticator) GenerateRefreshToken(claims jwt.Claims) (string, error) {
	return a.generateToken(claims, a.refreshSecret)
}

func (a *JWTAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(a.secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(a.issuer),
		jwt.WithIssuer(a.issuer),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
