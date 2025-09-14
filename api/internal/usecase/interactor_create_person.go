package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
)

type CreatePersonInteractor struct {
	personRepo domain.PersonRepository
}

func NewCreatePersonInteractor(personRepo domain.PersonRepository) *CreatePersonInteractor {
	return &CreatePersonInteractor{
		personRepo: personRepo,
	}
}

func (uc *CreatePersonInteractor) Execute(
	ctx context.Context,
	input CreatePersonInputData,
	presenter CreatePersonPresenter,
) error {
	emoji, err := domain.NewEmoji(input.Emoji)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	sign, err := domain.NewSign(input.Sign)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	coordinates, err := domain.NewCoordinates(input.Latitude, input.Longitude)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	parsedTime, err := time.Parse(time.RFC3339, input.SightingTime)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	sightingTime, err := domain.NewSightingTime(parsedTime)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	registerUUID, err := domain.NewUUID(input.RegisterUUID)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	gender, err := domain.NewGender(input.Gender)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	clothing, err := domain.NewClothing(input.Clothing)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	accessories, err := domain.NewAccessories(input.Accessories)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	vehicle, err := domain.NewVehicle(input.Vehicle)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	behavior, err := domain.NewBehavior(input.Behavior)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}
	hairstyle, err := domain.NewHairstyle(input.Hairstyle)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}

	person := domain.NewPerson(
		domain.GenerateUUID(),
		emoji,
		sign,
		registerUUID,
		0, // SightingCount 新規登録時は0
		sightingTime,
		coordinates,
		gender,
		clothing,
		accessories,
		vehicle,
		behavior,
		hairstyle,
	)
	if err := uc.personRepo.Create(ctx, &person); err != nil {
		if errors.Is(err, domain.ErrAlreadyExists) {
			return presenter.PresentError(fmt.Errorf("%w: %v", ErrBusinessRule, err))
		}
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrInternal, err))
	}

	output := CreatePersonOutputData{UUID: person.UUID().String()}
	return presenter.Present(output)
}
