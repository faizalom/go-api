package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/faizalom/go-api/internal/model"
	"github.com/faizalom/go-api/internal/service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_CreateUser(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	userHandler := NewUserHandler(mockUserService)

	reqBody := &model.NewUserRequest{
		Name:     "test user",
		Email:    "test@example.com",
		Password: "password",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	mockUserService.On("CreateUser", mock.Anything, mock.AnythingOfType("*model.NewUserRequest")).Return(&model.User{}, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.CreateUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserHandler_GetUserByID(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	userHandler := NewUserHandler(mockUserService)

	userID := uuid.New()
	req, err := http.NewRequest("GET", "/users/"+userID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", userID.String())

	mockUserService.On("GetUserByID", mock.Anything, userID).Return(&model.User{}, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.GetUserByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserHandler_UpdateUser(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	userHandler := NewUserHandler(mockUserService)

	userID := uuid.New()
	reqBody := &model.UpdateUserRequest{
		Name:  stringPtr("updated name"),
		Email: stringPtr("updated@example.com"),
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", "/users/"+userID.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", userID.String())

	mockUserService.On("UpdateUser", mock.Anything, userID, mock.AnythingOfType("*model.UpdateUserRequest")).Return(&model.User{}, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.UpdateUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserHandler_DeleteUser(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	userHandler := NewUserHandler(mockUserService)

	userID := uuid.New()
	req, err := http.NewRequest("DELETE", "/users/"+userID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", userID.String())

	mockUserService.On("DeleteUser", mock.Anything, userID).Return(nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.DeleteUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserHandler_ListUsers(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	userHandler := NewUserHandler(mockUserService)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockUserService.On("ListUsers", mock.Anything).Return([]*model.User{}, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ListUsers)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockUserService.AssertExpectations(t)
}

func stringPtr(s string) *string {
	return &s
}
