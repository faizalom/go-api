package router

import (
	"database/sql"
	"net/http"

	"github.com/faizalom/go-api/internal/handler"
	"github.com/faizalom/go-api/internal/repository"
	"github.com/faizalom/go-api/internal/service"
)

func NewUserRouter(db *sql.DB) http.Handler {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", userHandler.ListUsers)
	mux.HandleFunc("POST /", userHandler.CreateUser)
	mux.HandleFunc("GET /{id}", userHandler.GetUserByID)
	mux.HandleFunc("PUT /{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /{id}", userHandler.DeleteUser)
	return mux
}
