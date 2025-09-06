package usecase

import "github.com/iotassss/fushinsha-map-api/internal/domain"

type GetSuspiciousPersonDetailInputData struct {
	SuspiciousPersonUUID string
}

type SuspiciousPersonDetail struct {
	UUID          string  // UUID文字列
	Latitude      float64 // 座標値
	Longitude     float64
	Emoji         string
	Sign          string
	SightingCount int
	SightingTimes []string // ISO8601文字列
	Categories    []string
	Features      []string
	Gender        *string
	Clothing      *string
	Accessories   *string
	Vehicle       *string
	Behavior      *string
	Hairstyle     *string
}

type GetSuspiciousPersonDetailOutputData struct {
	SuspiciousPerson SuspiciousPersonDetail
}

type GetSuspiciousPersonDetailPresenter interface {
	Present(outputData GetSuspiciousPersonDetailOutputData) error
	PresentError(err error) error
}

type GetSuspiciousPersonDetailUsecase interface {
	Execute(
		input GetSuspiciousPersonDetailInputData,
		presenter GetSuspiciousPersonDetailPresenter,
		suspiciousPersonRepo domain.SuspiciousPersonRepository,
	) error
}
