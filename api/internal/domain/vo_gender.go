package domain

import "fmt"

type Gender string

var allowedGenders = map[string]struct{}{
	"男性": {},
	"女性": {},
	"不明": {},
}

func NewGender(s string) (Gender, error) {
	if _, ok := allowedGenders[s]; !ok {
		return "", fmt.Errorf("invalid gender: %s", s)
	}
	return Gender(s), nil
}

func (g Gender) String() string {
	return string(g)
}
