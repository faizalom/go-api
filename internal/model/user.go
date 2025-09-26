package model

import (
	"time"
	"github.com/google/uuid"
)

// User represents a user record in the database.
// This is the struct that will be returned in API responses.
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUserRequest defines the data required to create a new user.
type NewUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUserRequest defines the data allowed for updating a user.
// We use pointers to distinguish between a field not being provided
// and a field being provided with an empty value.
type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}
