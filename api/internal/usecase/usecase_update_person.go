package usecase

import (
	"context"
)

type UpdatePersonInputData struct {
	UUID         string
	Emoji        string
	Sign         string
	Categories   []string
	Features     []string
	SightingTime string // ISO8601
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
