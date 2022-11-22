package dao

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/domain"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/query"
)

func (receiver UserDao) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	var users []domain.User
	err := receiver.SelectMany(ctx, &users, query.C().Col("username").Eq(username).ToWhere())
	if err != nil {
		return false, err
	}
	if len(users) > 0 {
		return true, nil
	}
	return false, nil
}
