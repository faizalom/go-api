package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/faizalom/go-api/internal/config"
	"github.com/faizalom/go-api/internal/model"
	"github.com/faizalom/go-api/pkg/logger"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserClaimsKey contextKey = "userClaims"

// AuthMiddleware verifies the JWT token from the Authorization header.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			logger.Error.Println("Authorization header is missing")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 2. The header should be in the format "Bearer <token>"
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || strings.ToLower(headerParts[0]) != "bearer" {
			logger.Error.Println("Authorization header format is not Bearer {token}")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := headerParts[1]

		// 3. Parse and validate the token
		claims := &model.CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			// Use secret from config
			return []byte(config.App.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			logger.Error.Printf("Invalid token: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 4. Token is valid. Add claims to the context for downstream handlers
		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)

		// 5. Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
