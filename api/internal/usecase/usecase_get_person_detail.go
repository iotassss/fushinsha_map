package usecase

import "context"

type GetPersonDetailInputData struct {
	PersonUUID string
}

type PersonDetail struct {
	UUID          string   `json:"uuid"`     // UUID文字列
	Latitude      float64  `json:"latitude"` // 座標値
	Longitude     float64  `json:"longitude"`
	Emoji         string   `json:"emoji"`
	Sign          string   `json:"sign"`
	SightingCount int      `json:"sighting_count"`
	SightingTimes []string `json:"sighting_times"` // ISO8601文字列
	Categories    []string `json:"categories"`
	Gender        string   `json:"gender"`
	Clothing      string   `json:"clothing"`
	Accessories   string   `json:"accessories"`
	Vehicle       string   `json:"vehicle"`
	Behavior      string   `json:"behavior"`
	Hairstyle     string   `json:"hairstyle"`
}

type GetPersonDetailOutputData struct {
	Person PersonDetail
}

type GetPersonDetailPresenter interface {
	Present(outputData GetPersonDetailOutputData) error
	PresentError(err error) error
}

type GetPersonDetailUsecase interface {
	Execute(
		ctx context.Context,
		input GetPersonDetailInputData,
		presenter GetPersonDetailPresenter,
	) error
}
