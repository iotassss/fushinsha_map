package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type UUID struct {
	value uuid.UUID
}

func NewUUID(value string) (UUID, error) {
	parsedUUID, err := uuid.Parse(value)
	if err != nil {
		return UUID{}, fmt.Errorf("invalid UUID: %w", err)
	}
	return UUID{value: parsedUUID}, nil
}

func GenerateUUID() UUID {
	return UUID{value: uuid.New()}
}

func (id UUID) Value() uuid.UUID { return id.value }
func (id UUID) String() string   { return id.value.String() }
func (id UUID) IsNil() bool      { return id.value == uuid.Nil }
