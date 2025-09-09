package usecase

import "github.com/iotassss/fushinsha-map-api/internal/domain"

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
		input UpdatePersonInputData,
		presenter UpdatePersonPresenter,
		personRepo domain.PersonRepository,
	) error
}
