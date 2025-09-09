package usecase

import "github.com/iotassss/fushinsha-map-api/internal/domain"

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
		input IncrementSightingCountInputData,
		presenter IncrementSightingCountPresenter,
		personRepo domain.PersonRepository,
	) error
}
