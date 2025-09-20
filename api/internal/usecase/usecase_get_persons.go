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
	UUID          string  `json:"uuid"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Emoji         string  `json:"emoji"`
	Sign          string  `json:"sign"`
	SightingCount int     `json:"sighting_count"`
	SightingTime  string  `json:"sighting_time"` // ISO8601文字列
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
