package usecase

import (
	"context"
)

type FindOrCreateUserByIDTokenInput struct {
	IDToken string
}

type User struct {
	UUID string
}

type FindOrCreateUserByIDTokenPresenter interface {
	Present(user User) error
	PresentError(err error) error
}

type FindOrCreateUserByIDTokenUsecase interface {
	Execute(
		ctx context.Context,
		input FindOrCreateUserByIDTokenInput,
		presenter FindOrCreateUserByIDTokenPresenter,
	) error
}
