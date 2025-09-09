package usecase

import "context"

type LoginInputData struct {
	GoogleAccountID string
}

type LoginOutputData struct {
	Token string
}

type LoginPresenter interface {
	Present(outputData LoginOutputData) error
	PresentError(err error) error
}

type LoginUsecase interface {
	Execute(
		ctx context.Context,
		input LoginInputData,
		presenter LoginPresenter,
	) error
}
