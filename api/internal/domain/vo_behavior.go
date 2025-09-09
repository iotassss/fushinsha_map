package domain

import "fmt"

type Behavior string

var allowedBehaviors = map[string]struct{}{
	"":      {},
	"徘徊":    {},
	"大声":    {},
	"暴力":    {},
	"つきまとい": {},
	"その他":   {},
}

func NewBehavior(s string) (Behavior, error) {
	if _, ok := allowedBehaviors[s]; !ok {
		return "", fmt.Errorf("invalid behavior: %s", s)
	}
	return Behavior(s), nil
}

func (b Behavior) String() string {
	return string(b)
}
