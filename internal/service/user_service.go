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
	if _, err := s.repo.GetByEmail(ctx, req.Email); err == nil {
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
