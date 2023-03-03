package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("secReT@K3yyK4TanYYAAA4_")

type JWTClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
