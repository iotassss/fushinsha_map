package usecase

import "context"

type SignupInputData struct {
	GoogleAccountID string
}

type SignupOutputData struct {
	Success bool
}

type SignupPresenter interface {
	Present(outputData SignupOutputData) error
	PresentError(err error) error
}

type SignupUsecase interface {
	Execute(
		ctx context.Context,
		input SignupInputData,
		presenter SignupPresenter,
	) error
}
