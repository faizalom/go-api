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
	UpdateUser  http.HandlerFunc
	DeleteUser  http.HandlerFunc
	ListUsers   http.HandlerFunc
}

// New creates and configures a new router, injecting the handlers.
func New(h *Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	// Create a new router for the /api/v1 prefix
	apiV1Mux := http.NewServeMux()
	apiV1Mux.HandleFunc("/login", h.Login)
	apiV1Mux.Handle("/profile", protected(h.Profile))
	apiV1Mux.Handle("/example", protected(h.Example))

	// Mount the user router
	apiV1Mux.Handle("/users/", http.StripPrefix("/users", UserRouter(h)))


	// Wrap the apiV1Mux in a handler that strips the /api/v1 prefix
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1Mux))

	return mux
}

// protected is a helper that wraps a handler with standard protected-route middleware.
func protected(h http.Handler) http.Handler {
	return middleware.Chain(h, middleware.LoggingMiddleware, middleware.AuthMiddleware)
}