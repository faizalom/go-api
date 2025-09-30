package ierr

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with this email already exists")
	ErrUserNotFound      = errors.New("user not found")
)
