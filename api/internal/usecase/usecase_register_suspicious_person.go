package usecase

import "github.com/iotassss/fushinsha-map-api/internal/domain"

type RegisterSuspiciousPersonInputData struct {
	Latitude      float64
	Longitude     float64
	Emoji         string
	Sign          string
	Categories    []string
	Features      []string
	SightingTime  string // ISO8601
	RegistrarUUID string
}

type RegisterSuspiciousPersonOutputData struct {
	UUID string
}

type RegisterSuspiciousPersonPresenter interface {
	Present(outputData RegisterSuspiciousPersonOutputData) error
	PresentError(err error) error
}

type RegisterSuspiciousPersonUsecase interface {
	Execute(
		input RegisterSuspiciousPersonInputData,
		presenter RegisterSuspiciousPersonPresenter,
		suspiciousPersonRepo domain.SuspiciousPersonRepository,
	) error
}
