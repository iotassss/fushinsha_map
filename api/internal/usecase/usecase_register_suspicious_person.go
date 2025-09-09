package usecase

import "github.com/iotassss/fushinsha-map-api/internal/domain"

type RegisterPersonInputData struct {
	Latitude      float64
	Longitude     float64
	Emoji         string
	Sign          string
	Categories    []string
	Features      []string
	SightingTime  string // ISO8601
	RegistrarUUID string
}

type RegisterPersonOutputData struct {
	UUID string
}

type RegisterPersonPresenter interface {
	Present(outputData RegisterPersonOutputData) error
	PresentError(err error) error
}

type RegisterPersonUsecase interface {
	Execute(
		input RegisterPersonInputData,
		presenter RegisterPersonPresenter,
		personRepo domain.PersonRepository,
	) error
}
