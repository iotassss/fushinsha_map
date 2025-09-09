package domain

import "fmt"

type Accessories string

var allowedAccessories = map[string]struct{}{
	"":    {},
	"帽子":  {},
	"眼鏡":  {},
	"マスク": {},
	"バッグ": {},
	"なし":  {},
}

func NewAccessories(s string) (Accessories, error) {
	if _, ok := allowedAccessories[s]; !ok {
		return "", fmt.Errorf("invalid accessories: %s", s)
	}
	return Accessories(s), nil
}

func (a Accessories) String() string {
	return string(a)
}
