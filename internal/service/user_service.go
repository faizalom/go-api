package service

import (
	"context"
	"errors"

	"github.com/faizalom/go-api/internal/model"
	"github.com/faizalom/go-api/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists = errors.New("user with this email already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser handles the business logic for creating a new user.
func (s *UserService) CreateUser(ctx context.Context, req *model.NewUserRequest) (*model.User, error) {
	// Check if user already exists
	if _, _, err := s.repo.GetByEmail(ctx, req.Email); err == nil {
		return nil, ErrUserAlreadyExists
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Name:  req.Name,
		Email: req.Email,
	}

	// Call the repository to create the user
	createdUser, err := s.repo.Create(ctx, newUser, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		// In a real app, you'd check for sql.ErrNoRows and return a custom error
		return nil, ErrUserNotFound
	}
	return user, nil
}

// UpdateUser handles the business logic for updating a user.
func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, req *model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}

	if err := s.repo.Update(ctx, id, user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser handles the business logic for deleting a user.
func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return ErrUserNotFound
	}

	return s.repo.Delete(ctx, id)
}

// ListUsers retrieves a list of all users.
func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
