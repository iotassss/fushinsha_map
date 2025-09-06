package usecase

import "github.com/iotassss/fushinsha-map-api/internal/domain"

type UpdateSuspiciousPersonInputData struct {
	UUID         string
	Emoji        string
	Sign         string
	Categories   []string
	Features     []string
	SightingTime string // ISO8601
}

type UpdateSuspiciousPersonOutputData struct {
	Success bool
}

type UpdateSuspiciousPersonPresenter interface {
	Present(outputData UpdateSuspiciousPersonOutputData) error
	PresentError(err error) error
}

type UpdateSuspiciousPersonUsecase interface {
	Execute(
		input UpdateSuspiciousPersonInputData,
		presenter UpdateSuspiciousPersonPresenter,
		suspiciousPersonRepo domain.SuspiciousPersonRepository,
	) error
}
