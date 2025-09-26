package repository

import (
	"context"
	"database/sql"

	"github.com/faizalom/go-api/internal/model"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create inserts a new user record into the database.
func (r *UserRepository) Create(ctx context.Context, user *model.User, passwordHash string) (*model.User, error) {
	// SQL to insert user would go here.
	// It would return the generated ID, created_at, and updated_at.
	// For now, we'll simulate success.
	user.ID = uuid.New()
	// In a real scenario, these would be set by the database.
	// user.CreatedAt = time.Now()
	// user.UpdatedAt = time.Now()

	return user, nil
}

// GetByID retrieves a single user by their ID.
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	// SQL to SELECT a user by ID would go here.
	// For now, we'll return a nil user, indicating not found.
	return nil, sql.ErrNoRows
}

// GetByEmail retrieves a single user by their email.
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	// SQL to SELECT a user by email would go here.
	return nil, sql.ErrNoRows
}

// Update modifies an existing user record.
func (r *UserRepository) Update(ctx context.Context, id uuid.UUID, user *model.User) error {
	// SQL to UPDATE a user would go here.
	return nil
}

// Delete marks a user as deleted (soft delete).
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	// SQL to UPDATE the deleted_at timestamp would go here.
	return nil
}
