package domain

import "fmt"

type Coordinates struct {
	latitude  float64
	longitude float64
}

func NewCoordinates(lat, lng float64) (Coordinates, error) {
	if lat < -90.0 || lat > 90.0 {
		return Coordinates{}, fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng < -180.0 || lng > 180.0 {
		return Coordinates{}, fmt.Errorf("invalid longitude: %f", lng)
	}
	return Coordinates{latitude: lat, longitude: lng}, nil
}

func (c Coordinates) Latitude() float64 {
	return c.latitude
}

func (c Coordinates) Longitude() float64 {
	return c.longitude
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%f, %f)", c.latitude, c.longitude)
}
