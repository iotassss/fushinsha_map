package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
)

type GetSuspiciousPersonsInteractor struct {
	suspiciousPersonRepo domain.SuspiciousPersonRepository
}

func NewGetSuspiciousPersonsInteractor(suspiciousPersonRepo domain.SuspiciousPersonRepository) *GetSuspiciousPersonsInteractor {
	return &GetSuspiciousPersonsInteractor{
		suspiciousPersonRepo: suspiciousPersonRepo,
	}
}

func (uc *GetSuspiciousPersonsInteractor) Execute(
	ctx context.Context,
	input GetSuspiciousPersonsInputData,
	presenter GetSuspiciousPersonsPresenter,
) error {
	lx, err := strconv.ParseFloat(input.LX, 64)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	rx, err := strconv.ParseFloat(input.RX, 64)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	ty, err := strconv.ParseFloat(input.TY, 64)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	by, err := strconv.ParseFloat(input.BY, 64)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}

	area, err := domain.NewArea(lx, rx, ty, by)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrBusinessRule, err))
	}
	persons, err := uc.suspiciousPersonRepo.FindInArea(ctx, area)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrBusinessRule, err))
	}

	outputData := GetSuspiciousPersonsOutputData{}
	for _, p := range persons {
		summary := SuspiciousPersonSummary{
			UUID:          p.UUID().String(),
			Latitude:      p.Coordinates().Latitude(),
			Longitude:     p.Coordinates().Longitude(),
			Emoji:         p.Emoji().String(),
			Sign:          p.Sign().String(),
			SightingCount: p.SightingCount().Int(),
		}
		outputData.Persons = append(outputData.Persons, summary)
	}

	return presenter.Present(outputData)
}
