package router

import (
	"database/sql"

	"github.com/faizalom/go-api/internal/handler"
	"github.com/faizalom/go-api/internal/repository"
	"github.com/faizalom/go-api/internal/service"
)

func NewDependencies(db *sql.DB) *Handlers {
	// Repositories
	repoA := repository.NewRepoA(db)
	repoB := repository.NewRepoB(db)

	// Services
	serviceA := service.NewServiceA(repoA)
	serviceB := service.NewServiceB(repoB)

	// Handlers
	exampleHandler := handler.NewExampleHandler(serviceA, serviceB)

	// Assemble all handlers
	return &Handlers{
		Login:   handler.LoginHandler,
		Profile: handler.ProfileHandler,
		Example: exampleHandler.HandleRequest,
	}
}
