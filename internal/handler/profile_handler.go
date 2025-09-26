package handler

import (
	"fmt"
	"net/http"

	"github.com/faizalom/go-api/internal/middleware" // Our middleware package
	"github.com/faizalom/go-api/internal/model"
)

// ProfileHandler handles the user profile endpoint.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// The AuthMiddleware has already validated the token, so we can safely access the claims.
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*model.CustomClaims)
	if !ok {
		// This should not happen if the middleware is correctly applied.
		http.Error(w, "Internal Server Error: could not retrieve user claims", http.StatusInternalServerError)
		return
	}

	// Use the custom claims
	userID := claims.Subject
	userName := claims.Name
	userEmail := claims.Email
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Hello, %s (%s)", "user_id": "%s"}`, userName, userEmail, userID)
}
