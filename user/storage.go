package user

import "context"

type Storage interface {
	GetByEmail(context.Context, string) (*User, error)
	Add(context.Context, User) error
}
