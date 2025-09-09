package usecase

import "context"

type LogoutInputData struct {
	UserUUID string
}

type LogoutOutputData struct {
	Success bool
}

type LogoutPresenter interface {
	Present(outputData LogoutOutputData) error
	PresentError(err error) error
}

type LogoutUsecase interface {
	Execute(
		ctx context.Context,
		input LogoutInputData,
		presenter LogoutPresenter,
	) error
}
