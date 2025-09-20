package domain

import (
	"time"
)

type CreatedAt struct {
	time.Time
}

func NewCreatedAt(t time.Time) (CreatedAt, error) {
	if t.IsZero() {
		return CreatedAt{}, ErrValidation
	}
	return CreatedAt{t}, nil
}

func (c CreatedAt) String() string {
	return c.Time.Format(time.RFC3339)
}
