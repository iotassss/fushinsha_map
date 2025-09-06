package domain

import "fmt"

type Sign string

func NewSign(s string) (Sign, error) {
	runes := []rune(s)
	if len(runes) != 1 {
		return "", fmt.Errorf("invalid sign: only a single character is allowed")
	}
	return Sign(s), nil
}

func (s Sign) String() string {
	return string(s)
}
