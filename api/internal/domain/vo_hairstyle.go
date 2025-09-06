package domain

import "fmt"

type Hairstyle string

var allowedHairstyles = map[string]struct{}{
	"短髪":  {},
	"長髪":  {},
	"坊主":  {},
	"パーマ": {},
	"その他": {},
}

func NewHairstyle(s string) (Hairstyle, error) {
	if _, ok := allowedHairstyles[s]; !ok {
		return "", fmt.Errorf("invalid hairstyle: %s", s)
	}
	return Hairstyle(s), nil
}

func (h Hairstyle) String() string {
	return string(h)
}
