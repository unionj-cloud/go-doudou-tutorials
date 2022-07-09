package dao_test

import (
	"context"
	"ddldemo/dao"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserDaoImpl_FindUsersByHobby(t *testing.T) {
	t.Parallel()
	u := dao.NewUserDao(db)
	users, err := u.FindUsersByHobby(context.Background(), "football")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(users))
}
