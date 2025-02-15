package model

import "github.com/golang-jwt/jwt/v5"

type JwtClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
