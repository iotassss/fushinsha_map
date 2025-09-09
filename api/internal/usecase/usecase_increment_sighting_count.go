package usecase

import "context"

type IncrementSightingCountInputData struct {
	UserUUID   string
	PersonUUID string
}

type IncrementSightingCountOutputData struct {
	Success bool
}

type IncrementSightingCountPresenter interface {
	Present(outputData IncrementSightingCountOutputData) error
	PresentError(err error) error
}

type IncrementSightingCountUsecase interface {
	Execute(
		ctx context.Context,
		input IncrementSightingCountInputData,
		presenter IncrementSightingCountPresenter,
	) error
}
