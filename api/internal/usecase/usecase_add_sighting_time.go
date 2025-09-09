package usecase

import "github.com/iotassss/fushinsha-map-api/internal/domain"

type AddSightingTimeInputData struct {
	PersonUUID string
	Time       string // ISO8601形式
}

type AddSightingTimeOutputData struct {
	Success bool
}

type AddSightingTimePresenter interface {
	Present(outputData AddSightingTimeOutputData) error
	PresentError(err error) error
}

type AddSightingTimeUsecase interface {
	Execute(
		input AddSightingTimeInputData,
		presenter AddSightingTimePresenter,
		personRepo domain.PersonRepository,
	) error
}
