package router

import (
	"net/http"

	"github.com/faizalom/go-api/internal/middleware"
)

// Handlers now includes the user CRUD handlers.
type Handlers struct {
	Login       http.HandlerFunc
	Profile     http.HandlerFunc
	Example     http.HandlerFunc
	CreateUser  http.HandlerFunc
	GetUserByID http.HandlerFunc
}

// New creates and configures a new router, injecting the handlers.
func New(h *Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/login", h.Login)
	mux.HandleFunc("POST /users", h.CreateUser) // Create a new user

	// For this route, we are using the new enhanced routing patterns from Go 1.22+
	// The {id} part is a wildcard that can be accessed in the handler.
	mux.HandleFunc("GET /api/v1/user/{id}", h.GetUserByID)

	// Protected routes
	mux.Handle("/api/v1/profile", protected(h.Profile))
	mux.Handle("/example", protected(h.Example))

	return mux
}

// protected is a helper that wraps a handler with standard protected-route middleware.
func protected(h http.Handler) http.Handler {
	return middleware.Chain(h, middleware.LoggingMiddleware, middleware.AuthMiddleware)
}
