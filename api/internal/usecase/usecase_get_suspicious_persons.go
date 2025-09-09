package usecase

import (
	"context"
)

type GetPersonsInputData struct {
	LX string // 左上X
	RX string // 右下X
	TY string // 左上Y
	BY string // 右下Y
}

type PersonSummary struct {
	UUID          string
	Latitude      float64
	Longitude     float64
	Emoji         string
	Sign          string
	SightingCount int
}

type GetPersonsOutputData struct {
	Persons []PersonSummary
}

type GetPersonsPresenter interface {
	Present(outputData GetPersonsOutputData) error
	PresentError(err error) error
}

type GetPersonsUsecase interface {
	Execute(
		ctx context.Context,
		input GetPersonsInputData,
		presenter GetPersonsPresenter,
	) error
}
