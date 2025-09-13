package domain

import "fmt"

type Clothing string

var allowedClothings = map[string]struct{}{
	"":    {},
	"スーツ": {},
	"制服":  {},
	"私服":  {},
	"作業着": {},
	"その他": {},
}

func NewClothing(s string) (Clothing, error) {
	if _, ok := allowedClothings[s]; !ok {
		return "", fmt.Errorf("invalid clothing: %s", s)
	}
	return Clothing(s), nil
}

func (c Clothing) String() string {
	return string(c)
}
