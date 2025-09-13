package usecase

import "context"

type GetPersonDetailInputData struct {
	PersonUUID string
}

type PersonDetail struct {
	UUID          string  // UUID文字列
	Latitude      float64 // 座標値
	Longitude     float64
	Emoji         string
	Sign          string
	SightingCount int
	SightingTimes []string // ISO8601文字列
	Categories    []string
	Features      []string
	Gender        string
	Clothing      string
	Accessories   string
	Vehicle       string
	Behavior      string
	Hairstyle     string
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
