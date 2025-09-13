package domain

import "errors"

var (
	ErrValidation    = errors.New("validation error")
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrRepository    = errors.New("repository error")
)
