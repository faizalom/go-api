package mocks

import (
	"context"

	"github.com/faizalom/go-api/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, req *model.NewUserRequest) (*model.User, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, id uuid.UUID, req *model.UpdateUserRequest) (*model.User, error) {
	args := m.Called(ctx, id, req)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.User), args.Error(1)
}
