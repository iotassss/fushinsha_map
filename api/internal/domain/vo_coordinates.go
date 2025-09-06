package domain

import "fmt"

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func NewCoordinates(lat, lng float64) (Coordinates, error) {
	if lat < -90.0 || lat > 90.0 {
		return Coordinates{}, fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng < -180.0 || lng > 180.0 {
		return Coordinates{}, fmt.Errorf("invalid longitude: %f", lng)
	}
	return Coordinates{Latitude: lat, Longitude: lng}, nil
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%f, %f)", c.Latitude, c.Longitude)
}
