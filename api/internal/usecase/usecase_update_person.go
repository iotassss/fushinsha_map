package usecase

import (
	"context"
)

type UpdatePersonInputData struct {
	UUID          string
	SightingCount int
}

type UpdatePersonOutputData struct {
	Success bool
}

type UpdatePersonPresenter interface {
	Present(outputData UpdatePersonOutputData) error
	PresentError(err error) error
}

type UpdatePersonUsecase interface {
	Execute(
		ctx context.Context,
		input UpdatePersonInputData,
		presenter UpdatePersonPresenter,
	) error
}
