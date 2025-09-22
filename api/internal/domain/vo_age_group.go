package domain

import "fmt"

type AgeGroup string

var allowedAgeGroups = map[string]struct{}{
	"未成年":   {},
	"20代":   {},
	"30代":   {},
	"40代":   {},
	"50代":   {},
	"60代以上": {},
}

func NewAgeGroup(s string) (AgeGroup, error) {
	if _, ok := allowedAgeGroups[s]; !ok {
		return "", fmt.Errorf("invalid age group: %s", s)
	}
	return AgeGroup(s), nil
}

func (ag AgeGroup) String() string {
	return string(ag)
}
