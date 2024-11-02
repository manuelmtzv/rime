package auth

import "github.com/golang-jwt/jwt/v5"

type Authenticator interface {
	GenerateAccessToken(claims jwt.Claims) (string, error)
	GenerateRefreshToken(claims jwt.Claims) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
