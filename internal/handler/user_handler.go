package handler

import (
	"encoding/json"
	"net/http"

	"github.com/faizalom/go-api/internal/model"
	"github.com/faizalom/go-api/internal/service"

	"github.com/google/uuid"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// CreateUser handles the HTTP request for creating a new user.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req model.NewUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// In a real app, you would add validation for the request here.

	createdUser, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		// A more robust implementation would check the type of error
		if err == service.ErrUserAlreadyExists {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// GetUserByID handles the HTTP request for retrieving a user by their ID.
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// In a real app, you would get the ID from the URL path, e.g., /users/{id}
	// For this example, we'll imagine we have the ID.
	id, _ := uuid.NewRandom() // Placeholder

	user, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		if err == service.ErrUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
