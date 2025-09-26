package handler

import (
	"net/http"
	"time"

	"github.com/faizalom/go-api/internal/config"
	"github.com/faizalom/go-api/internal/model" // Our new model package
	"github.com/faizalom/go-api/pkg/logger"

	"github.com/golang-jwt/jwt/v5"
)

// LoginHandler simulates a user login and generates a JWT.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// In a real application, you would validate username/password here.
	// For this example, we'll assume authentication is successful.

	// Create the custom claims
	claims := model.CustomClaims{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		RegisteredClaims: jwt.RegisteredClaims{
			// Use the user's ID as the subject
			Subject: "user-12345",
			// Set token expiration
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret from config
	tokenString, err := token.SignedString([]byte(config.App.JWT.Secret))
	if err != nil {
		logger.Error.Printf("Could not sign token: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the token back to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + tokenString + `"}`))
}
