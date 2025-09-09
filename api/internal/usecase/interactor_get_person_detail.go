package usecase

import (
	"context"
	"fmt"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
)

type GetPersonDetailInteractor struct {
	personRepo domain.PersonRepository
}

func NewGetPersonDetailInteractor(personRepo domain.PersonRepository) *GetPersonDetailInteractor {
	return &GetPersonDetailInteractor{
		personRepo: personRepo,
	}
}

func (uc *GetPersonDetailInteractor) Execute(
	ctx context.Context,
	input GetPersonDetailInputData,
	presenter GetPersonDetailPresenter,
) error {
	uuid, err := domain.NewUUID(input.PersonUUID)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	person, err := uc.personRepo.FindByUUID(ctx, uuid)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrNotFound, err))
	}

	output := GetPersonDetailOutputData{
		Person: PersonDetail{
			UUID:          person.UUID().String(),
			Latitude:      person.Coordinates().Latitude(),
			Longitude:     person.Coordinates().Longitude(),
			Emoji:         person.Emoji().String(),
			Sign:          person.Sign().String(),
			SightingCount: person.SightingCount().Int(),
			Gender:        toStringPtr(person.Gender()),
			Clothing:      toStringPtr(person.Clothing()),
			Accessories:   toStringPtr(person.Accessories()),
			Vehicle:       toStringPtr(person.Vehicle()),
			Behavior:      toStringPtr(person.Behavior()),
			Hairstyle:     toStringPtr(person.Hairstyle()),
		},
	}

	return presenter.Present(output)
}

func toStringPtr[T any](v *T) *string {
	if v == nil {
		return nil
	}
	s := fmt.Sprintf("%v", *v)
	return &s
}
