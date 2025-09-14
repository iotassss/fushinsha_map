package usecase

import "context"

type CreatePersonInputData struct {
	Latitude     float64
	Longitude    float64
	Emoji        string
	Sign         string
	Categories   []string
	Features     []string
	SightingTime string // ISO8601
	RegisterUUID string
	Gender       string
	Clothing     string
	Accessories  string
	Vehicle      string
	Behavior     string
	Hairstyle    string
}

type CreatePersonOutputData struct {
	UUID string
}

type CreatePersonPresenter interface {
	Present(outputData CreatePersonOutputData) error
	PresentError(err error) error
}

type CreatePersonUsecase interface {
	Execute(
		ctx context.Context,
		input CreatePersonInputData,
		presenter CreatePersonPresenter,
	) error
}
