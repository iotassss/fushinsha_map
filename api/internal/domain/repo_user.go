package domain

import "context"

type UserRepository interface {
	FindByGoogleAccountID(ctx context.Context, googleAccountID GoogleAccountID) (*User, error)
	Create(ctx context.Context, user *User) error
}
