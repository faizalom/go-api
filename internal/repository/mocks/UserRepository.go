package mocks

import (
	"context"

	"github.com/faizalom/go-api/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *model.User, passwordHash string) (*model.User, error) {
	args := m.Called(ctx, user, passwordHash)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, string, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*model.User), args.String(1), args.Error(2)
}

func (m *MockUserRepository) Update(ctx context.Context, id uuid.UUID, user *model.User) error {
	args := m.Called(ctx, id, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) List(ctx context.Context) ([]*model.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.User), args.Error(1)
}
