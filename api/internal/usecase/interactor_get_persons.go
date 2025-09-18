package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
)

type GetPersonsInteractor struct {
	personRepo domain.PersonRepository
}

const maxPersonsInArea = 100

func NewGetPersonsInteractor(personRepo domain.PersonRepository) *GetPersonsInteractor {
	return &GetPersonsInteractor{
		personRepo: personRepo,
	}
}

func (uc *GetPersonsInteractor) Execute(
	ctx context.Context,
	input GetPersonsInputData,
	presenter GetPersonsPresenter,
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
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	persons, err := uc.personRepo.FindInArea(ctx, area, maxPersonsInArea)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrInternal, err))
	}

	outputData := GetPersonsOutputData{}
	for _, p := range persons {
		summary := PersonSummary{
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
