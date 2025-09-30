package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faizalom/go-api/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	now := time.Now()
	user := &model.User{
		Name:  "test user",
		Email: "test@example.com",
	}
	passwordHash := "password_hash"
	newUUID := uuid.New()

	mock.ExpectQuery(`INSERT INTO users`).
		WithArgs(user.Name, user.Email, passwordHash).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
			AddRow(newUUID, now, now))

	createdUser, err := repo.Create(context.Background(), user, passwordHash)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, newUUID, createdUser.ID)
	assert.Equal(t, now, createdUser.CreatedAt)
	assert.Equal(t, now, createdUser.UpdatedAt)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	now := time.Now()
	user := &model.User{
		ID:        uuid.New(),
		Name:      "test user",
		Email:     "test@example.com",
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "is_active", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.IsActive, user.CreatedAt, user.UpdatedAt)

	mock.ExpectQuery(`SELECT id, name, email, is_active, created_at, updated_at FROM users WHERE id = \$1`).
		WithArgs(user.ID).
		WillReturnRows(rows)

	foundUser, err := repo.GetByID(context.Background(), user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user, foundUser)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	now := time.Now()
	user := &model.User{
		ID:        uuid.New(),
		Name:      "test user",
		Email:     "test@example.com",
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	passwordHash := "password_hash"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password_hash", "is_active", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, passwordHash, user.IsActive, user.CreatedAt, user.UpdatedAt)

	mock.ExpectQuery(`SELECT id, name, email, password_hash, is_active, created_at, updated_at FROM users WHERE email = \$1`).
		WithArgs(user.Email).
		WillReturnRows(rows)

	foundUser, foundPasswordHash, err := repo.GetByEmail(context.Background(), user.Email)

	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user, foundUser)
	assert.Equal(t, passwordHash, foundPasswordHash)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	user := &model.User{
		ID:    uuid.New(),
		Name:  "updated name",
		Email: "updated@example.com",
	}

	mock.ExpectExec(`UPDATE users`).
		WithArgs(user.Name, user.Email, user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(context.Background(), user.ID, user)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	userID := uuid.New()

	mock.ExpectExec(`UPDATE users`).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(context.Background(), userID)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	now := time.Now()
	users := []*model.User{
		{
			ID:        uuid.New(),
			Name:      "test user 1",
			Email:     "test1@example.com",
			IsActive:  true,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        uuid.New(),
			Name:      "test user 2",
			Email:     "test2@example.com",
			IsActive:  true,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "is_active", "created_at", "updated_at"})
	for _, user := range users {
		rows.AddRow(user.ID, user.Name, user.Email, user.IsActive, user.CreatedAt, user.UpdatedAt)
	}

	mock.ExpectQuery(`SELECT id, name, email, is_active, created_at, updated_at FROM users`).
		WillReturnRows(rows)

	foundUsers, err := repo.List(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, foundUsers)
	assert.Equal(t, users, foundUsers)
	assert.NoError(t, mock.ExpectationsWereMet())
}
