package auth

import "github.com/golang-jwt/jwt/v5"

type Authenticator interface {
	GenerateToken(claims jwt.Claims, tokenType string) (string, error)
	ValidateToken(token string, tokenType string) (*jwt.Token, error)
}
