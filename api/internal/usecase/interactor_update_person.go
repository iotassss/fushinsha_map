package usecase

import (
	"context"
	"fmt"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
)

type UpdatePersonInteractor struct {
	personRepo domain.PersonRepository
}

func NewUpdatePersonInteractor(personRepo domain.PersonRepository) *UpdatePersonInteractor {
	return &UpdatePersonInteractor{
		personRepo: personRepo,
	}
}

func (uc *UpdatePersonInteractor) Execute(
	ctx context.Context,
	input UpdatePersonInputData,
	presenter UpdatePersonPresenter,
) error {
	uuid, err := domain.NewUUID(input.UUID)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	sightingCount, err := domain.NewSightingCount(input.SightingCount)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}

	person, err := uc.personRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrNotFound, err))
	}

	person.SetSightingCount(sightingCount)

	if err := uc.personRepo.Update(ctx, person); err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrInternal, err))
	}

	return presenter.Present(UpdatePersonOutputData{Success: true})
}
