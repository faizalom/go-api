package service

import (
	"context"
	"testing"

	"github.com/faizalom/go-api/internal/model"
	"github.com/faizalom/go-api/internal/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreateUser(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	req := &model.NewUserRequest{
		Name:     "test user",
		Email:    "test@example.com",
		Password: "password",
	}

	mockUserRepo.On("GetByEmail", mock.Anything, req.Email).Return(&model.User{}, "", assert.AnError)
	mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.User"), mock.AnythingOfType("string")).Return(&model.User{}, nil)

	createdUser, err := userService.CreateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	user := &model.User{
		ID: uuid.New(),
	}

	mockUserRepo.On("GetByID", mock.Anything, user.ID).Return(user, nil)

	foundUser, err := userService.GetUserByID(context.Background(), user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user, foundUser)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	userID := uuid.New()
	req := &model.UpdateUserRequest{
		Name:  stringPtr("updated name"),
		Email: stringPtr("updated@example.com"),
	}

	user := &model.User{
		ID:    userID,
		Name:  "original name",
		Email: "original@example.com",
	}

	mockUserRepo.On("GetByID", mock.Anything, userID).Return(user, nil)
	mockUserRepo.On("Update", mock.Anything, userID, mock.AnythingOfType("*model.User")).Return(nil)

	updatedUser, err := userService.UpdateUser(context.Background(), userID, req)

	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, *req.Name, updatedUser.Name)
	assert.Equal(t, *req.Email, updatedUser.Email)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	userID := uuid.New()

	mockUserRepo.On("GetByID", mock.Anything, userID).Return(&model.User{}, nil)
	mockUserRepo.On("Delete", mock.Anything, userID).Return(nil)

	err := userService.DeleteUser(context.Background(), userID)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_ListUsers(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	users := []*model.User{
		{ID: uuid.New()},
		{ID: uuid.New()},
	}

	mockUserRepo.On("List", mock.Anything).Return(users, nil)

	foundUsers, err := userService.ListUsers(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, foundUsers)
	assert.Equal(t, users, foundUsers)
	mockUserRepo.AssertExpectations(t)
}

func stringPtr(s string) *string {
	return &s
}
