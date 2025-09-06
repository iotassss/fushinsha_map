package domain

import "fmt"

type SightingCount int

func NewSightingCount(n int) (SightingCount, error) {
	if n < 0 {
		return 0, fmt.Errorf("invalid sighting count: must be >= 0")
	}
	return SightingCount(n), nil
}

func (sc SightingCount) Int() int {
	return int(sc)
}
