package repository

import (
	"context"
	"database/sql"

	"github.com/faizalom/go-api/internal/ierr"
	"github.com/faizalom/go-api/internal/model"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{DB: db}
}

// Create inserts a new user record into the database.
func (r *UserRepository) Create(ctx context.Context, user *model.User, passwordHash string) (*model.User, error) {
	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	err := r.DB.QueryRowContext(ctx, query, user.Name, user.Email, passwordHash).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetByID retrieves a single user by their ID.
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	query := `
		SELECT id, name, email, is_active, created_at, updated_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`
	user := &model.User{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ierr.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

// GetByEmail retrieves a single user by their email.
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, string, error) {
	query := `
		SELECT id, name, email, password_hash, is_active, created_at, updated_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`
	user := &model.User{}
	var passwordHash string
	err := r.DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &passwordHash, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", ierr.ErrUserNotFound
		}
		return nil, "", err
	}
	return user, passwordHash, nil
}

// Update modifies an existing user record.
func (r *UserRepository) Update(ctx context.Context, id uuid.UUID, user *model.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, updated_at = NOW()
		WHERE id = $3 AND deleted_at IS NULL
	`
	_, err := r.DB.ExecContext(ctx, query, user.Name, user.Email, id)
	return err
}

// Delete marks a user as deleted (soft delete).
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE users
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

// List retrieves a list of users from the database.
func (r *UserRepository) List(ctx context.Context) ([]*model.User, error) {
	query := `
		SELECT id, name, email, is_active, created_at, updated_at
		FROM users
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}