package main

import (
	"database/sql"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver

	"github.com/faizalom/go-api/internal/config"
	"github.com/faizalom/go-api/internal/handler"
	"github.com/faizalom/go-api/internal/repository"
	"github.com/faizalom/go-api/internal/router"
	"github.com/faizalom/go-api/internal/service"
	"github.com/faizalom/go-api/pkg/logger"
)

func main() {
	logger.Init()

	if err := config.Load("../../configs/config.local.yaml"); err != nil {
		logger.Error.Fatalf("Could not load configuration: %v", err)
	}

	//============================================================================
	// Database Connection
	//============================================================================

	db, err := sql.Open("pgx", config.App.Database.DSN)
	if err != nil {
		logger.Error.Fatalf("Could not open database connection: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		logger.Error.Fatalf("Could not ping database: %v", err)
	}
	logger.Info.Println("Successfully connected to the database.")

	//============================================================================
	// Dependency Injection
	//============================================================================

	// Repositories
	repoA := repository.NewRepoA(db)
	repoB := repository.NewRepoB(db)
	userRepo := repository.NewUserRepository(db)

	// Services
	serviceA := service.NewServiceA(repoA)
	serviceB := service.NewServiceB(repoB)
	userService := service.NewUserService(userRepo)

	// Handlers
	exampleHandler := handler.NewExampleHandler(serviceA, serviceB)
	userHandler := handler.NewUserHandler(userService)

	// Assemble all handlers
	h := &router.Handlers{
		Login:       handler.LoginHandler,
		Profile:     handler.ProfileHandler,
		Example:     exampleHandler.HandleRequest,
		CreateUser:  userHandler.CreateUser,
		GetUserByID: userHandler.GetUserByID,
	}

	//============================================================================

	logger.Info.Println("Starting the workout API server...")

	r := router.New(h)
	addr := config.App.Server.Port
	logger.Info.Printf("Server is listening on http://localhost%s", addr)

	err = http.ListenAndServe(addr, r)
	if err != nil {
		logger.Error.Fatalf("Could not start server: %s\n", err)
	}
}
