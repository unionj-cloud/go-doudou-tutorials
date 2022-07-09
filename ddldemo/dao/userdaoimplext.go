package dao

import (
	"context"
	"ddldemo/domain"
)

func (receiver UserDaoImpl) FindUsersByHobby(ctx context.Context, hobby string) (users []domain.User, err error) {
	sqlStr := `select * from ddl_user where hobby = ? and delete_at is null`
	err = receiver.db.SelectContext(ctx, &users, receiver.db.Rebind(sqlStr), hobby)
	return
}
