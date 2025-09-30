package repository

import (
	"context"

	"github.com/faizalom/go-api/internal/model"

	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User, passwordHash string) (*model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, string, error)
	Update(ctx context.Context, id uuid.UUID, user *model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]*model.User, error)
}
