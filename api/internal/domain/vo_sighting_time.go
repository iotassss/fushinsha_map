package domain

import (
	"fmt"
	"time"
)

type SightingTime struct {
	time time.Time
}

func NewSightingTime(t time.Time) (SightingTime, error) {
	if t.IsZero() {
		return SightingTime{}, fmt.Errorf("invalid sighting time: zero value is not allowed")
	}
	return SightingTime{time: t}, nil
}

func (st SightingTime) Time() time.Time {
	return st.time
}

func (st SightingTime) String() string {
	return st.time.Format("15:04")
}
