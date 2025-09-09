package usecase

import (
	"context"
)

type GetSuspiciousPersonsInputData struct {
	LX string // 左上X
	RX string // 右下X
	TY string // 左上Y
	BY string // 右下Y
}

type SuspiciousPersonSummary struct {
	UUID          string
	Latitude      float64
	Longitude     float64
	Emoji         string
	Sign          string
	SightingCount int
}

type GetSuspiciousPersonsOutputData struct {
	Persons []SuspiciousPersonSummary
}

type GetSuspiciousPersonsPresenter interface {
	Present(outputData GetSuspiciousPersonsOutputData) error
	PresentError(err error) error
}

type GetSuspiciousPersonsUsecase interface {
	Execute(
		ctx context.Context,
		input GetSuspiciousPersonsInputData,
		presenter GetSuspiciousPersonsPresenter,
	) error
}
