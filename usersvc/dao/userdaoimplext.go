package dao

import (
	"context"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/query"
	"usersvc/domain"
)

func (receiver UserDaoImpl) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	many, err := receiver.SelectMany(ctx, query.C().Col("username").Eq(username))
	if err != nil {
		return false, err
	}
	users := many.([]domain.User)
	if len(users) > 0 {
		return true, nil
	}
	return false, nil
}
