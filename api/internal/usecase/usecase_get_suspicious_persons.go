package usecase

import "github.com/iotassss/fushinsha-map-api/internal/domain"

type GetSuspiciousPersonsInputData struct {
	LX float64 // 左上X
	RX float64 // 右下X
	TY float64 // 左上Y
	BY float64 // 右下Y
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
		input GetSuspiciousPersonsInputData,
		presenter GetSuspiciousPersonsPresenter,
		suspiciousPersonRepo domain.SuspiciousPersonRepository,
	) error
}
