package dao

import "context"

type UserDao interface {
	Base
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
}
