package domain

import "fmt"

type Vehicle string

var allowedVehicles = map[string]struct{}{
	"自転車": {},
	"バイク": {},
	"自動車": {},
	"徒歩":  {},
	"その他": {},
}

func NewVehicle(s string) (Vehicle, error) {
	if _, ok := allowedVehicles[s]; !ok {
		return "", fmt.Errorf("invalid vehicle: %s", s)
	}
	return Vehicle(s), nil
}

func (v Vehicle) String() string {
	return string(v)
}
