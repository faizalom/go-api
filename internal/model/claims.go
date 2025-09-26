package model

import "github.com/golang-jwt/jwt/v5"

// CustomClaims contains our custom data and the standard registered claims.
type CustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
