package domain

import (
	"fmt"
)

type GoogleAccountID struct {
	value string
}

func NewGoogleAccountID(value string) (GoogleAccountID, error) {
	if value == "" {
		return GoogleAccountID{}, fmt.Errorf("GoogleAccountID cannot be empty")
	}
	return GoogleAccountID{value: value}, nil
}

func (id GoogleAccountID) Value() string { return id.value }
func (id GoogleAccountID) String() string { return id.value }
