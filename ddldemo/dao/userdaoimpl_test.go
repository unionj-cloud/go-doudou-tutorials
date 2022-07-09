package dao_test

import (
	"context"
	"ddldemo/dao"
	"ddldemo/domain"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	_ "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/logger"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/query"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/sortenum"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/wrapper"
	"log"
	"os"
	"testing"
	"time"
)

type DbConfig struct {
	Host    string
	Port    string
	User    string
	Passwd  string
	Schema  string
	Charset string
}

var db *sqlx.DB

func setUp() (int, func()) {
	u := dao.NewUserDao(db)
	school := "Shanxi Univ."
	user := domain.User{
		Name:   "jack",
		Phone:  "13552053960",
		Age:    30,
		School: &school,
		Hobby:  "football",
	}
	_, err := u.Insert(context.Background(), &user)
	if err != nil {
		panic(err)
	}
	return user.Id, func() {
		u := dao.NewUserDao(db)
		_, err = u.DeleteMany(context.Background(), query.C().Col("name").Ne(""))
		if err != nil {
			panic(err)
		}
	}
}

func TestMain(m *testing.M) {
	os.Setenv("GDD_SQL_LOG_ENABLE", "true")
	defer os.Unsetenv("GDD_SQL_LOG_ENABLE")
	var dbConfig DbConfig
	err := envconfig.Process("db", &dbConfig)
	if err != nil {
		logrus.Fatal("Error processing env", err)
	}
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbConfig.User,
		dbConfig.Passwd,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Schema,
		dbConfig.Charset)
	conn += `&loc=Asia%2FShanghai&parseTime=True`
	db, err = sqlx.Connect("mysql", conn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	db.MapperFunc(strcase.ToSnake)
	m.Run()
}

// test that non-zero column should be updated to zero value
func TestUserDaoImpl_UpdateMany(t *testing.T) {
	t.Parallel()
	userId, tearDown := setUp()
	defer tearDown()
	gdddb := wrapper.NewGddDB(db, wrapper.WithLogger(logger.NewSqlLogger(log.Default())))
	// begin transaction
	tx, err := gdddb.BeginTxx(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			_ = tx.Rollback()
			t.Fatal(err)
		}
	}()

	u := dao.NewUserDao(tx)
	q := `select * from ddl_user where id=? for update`
	var user domain.User
	err = tx.GetContext(context.Background(), &user, tx.Rebind(q), userId)
	if err != nil {
		panic(err)
	}
	require.NotZero(t, user.School)

	avg, _ := decimal.NewFromString("95.88")
	user = domain.User{
		Name:     "jack",
		Age:      10,
		AvgScore: avg,
		Hobby:    "football",
	}

	where := query.C().Col("school").Eq("Shanxi Univ.").
		And(query.C().Col("delete_at").IsNull())

	got, err := u.UpdateMany(context.Background(), user, where)
	if err != nil {
		panic(err)
	}
	require.NotEqual(t, 0, got)

	get, err := u.Get(context.Background(), userId)
	if err != nil {
		panic(err)
	}
	user = get.(domain.User)
	require.Zero(t, user.School)

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func TestUserDaoImpl_UpsertNoneZero(t *testing.T) {
	t.Parallel()
	gdddb := wrapper.NewGddDB(db, wrapper.WithLogger(logger.NewSqlLogger(log.Default())))
	u := dao.NewUserDao(gdddb)
	avg, _ := decimal.NewFromString("85.88")
	now := time.Now().Local()
	user := domain.User{
		AvgScore:  avg,
		IsStudent: 1,
		Name:      "David3",
		DeleteAt:  &now,
		Hobby:     "football",
	}

	got, err := u.UpsertNoneZero(context.Background(), user)
	if err != nil {
		panic(err)
	}
	require.NotEqual(t, 0, got)
}

func TestUserDaoImpl_DeleteMany(t *testing.T) {
	t.Parallel()
	userId, tearDown := setUp()
	defer tearDown()
	gdddb := wrapper.NewGddDB(db, wrapper.WithLogger(logger.NewSqlLogger(log.Default())))
	// begin transaction
	tx, err := gdddb.BeginTxx(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	u := dao.NewUserDao(tx)
	where := query.C().Col("delete_at").IsNotNull()
	got, err := u.DeleteMany(context.Background(), where)
	if err != nil {
		panic(err)
	}
	require.NotEqual(t, 0, got)

	school := "Shanxi Univ."
	user := domain.User{
		Id:     userId,
		Name:   "jack",
		Phone:  "13552053960",
		Age:    30,
		School: &school,
		Hobby:  "football",
	}
	_, err = u.Insert(context.Background(), &user)
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

// test that non-zero column should not be updated to zero value
func TestUserDaoImpl_UpdateNoneZero(t *testing.T) {
	t.Parallel()
	userId, tearDown := setUp()
	defer tearDown()

	gdddb := wrapper.NewGddDB(db, wrapper.WithLogger(logger.NewSqlLogger(log.Default())))
	u := dao.NewUserDao(gdddb)
	get, err := u.Get(context.Background(), userId)
	if err != nil {
		panic(err)
	}
	user := get.(domain.User)
	require.NotZero(t, user.Age)

	avg, _ := decimal.NewFromString("85.88")
	now := time.Now().Local()
	user = domain.User{
		Id:        userId,
		AvgScore:  avg,
		IsStudent: 1,
		Name:      "David3",
		DeleteAt:  &now,
		Hobby:     "football",
	}

	got, err := u.UpsertNoneZero(context.Background(), user)
	if err != nil {
		panic(err)
	}
	require.NotEqual(t, 0, got)

	get, err = u.Get(context.Background(), userId)
	if err != nil {
		panic(err)
	}
	user = get.(domain.User)
	require.NotZero(t, user.Age)
}

func TestUserDaoImpl_UpdateManyNoneZero(t *testing.T) {
	t.Parallel()
	_, tearDown := setUp()
	defer tearDown()
	gdddb := wrapper.NewGddDB(db, wrapper.WithLogger(logger.NewSqlLogger(log.Default())))
	// begin transaction
	tx, err := gdddb.BeginTxx(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			_ = tx.Rollback()
			t.Fatal(err)
		}
	}()

	u := dao.NewUserDao(tx)
	avg, _ := decimal.NewFromString("95.88")
	user := domain.User{
		Age:      10,
		AvgScore: avg,
		Hobby:    "football",
	}

	where := query.C().Col("school").Eq("Shanxi Univ.").
		And(query.C().Col("delete_at").IsNull())

	got, err := u.UpdateManyNoneZero(context.Background(), user, where)
	if err != nil {
		panic(err)
	}
	require.NotEqual(t, 0, got)

	where = query.C().Col("school").Eq("Shanxi Univ.").
		And(query.C().Col("delete_at").IsNull())

	many, err := u.SelectMany(context.Background(), where)
	if err != nil {
		panic(err)
	}
	users := many.([]domain.User)
	for _, item := range users {
		require.NotZero(t, item.School)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func TestUserDaoImpl_CountMany(t *testing.T) {
	t.Parallel()
	_, tearDown := setUp()
	defer tearDown()
	gdddb := wrapper.NewGddDB(db, wrapper.WithLogger(logger.NewSqlLogger(log.Default())))
	u := dao.NewUserDao(gdddb)

	where := query.C().Col("school").Eq("Shanxi Univ.").
		And(query.C().Col("delete_at").IsNull())

	got, err := u.CountMany(context.Background(), where)
	if err != nil {
		panic(err)
	}
	require.NotEqual(t, 0, got)
}

func TestUserDaoImpl_PageMany(t *testing.T) {
	t.Parallel()
	_, tearDown := setUp()
	defer tearDown()
	gdddb := wrapper.NewGddDB(db, wrapper.WithLogger(logger.NewSqlLogger(log.Default())))
	u := dao.NewUserDao(gdddb)

	where := query.C().Col("school").Eq("Shanxi Univ.").
		And(query.C().Col("delete_at").IsNull())

	page := query.Page{
		Orders: []query.Order{
			{
				Col:  "create_at",
				Sort: sortenum.Desc,
			},
		},
		Offset: 0,
		Size:   2,
	}
	got, err := u.PageMany(context.Background(), page, where)
	if err != nil {
		panic(err)
	}
	require.NotEqual(t, 0, len(got.Items.([]domain.User)))
}

func TestUserDaoImpl_DeleteManySoft(t *testing.T) {
	t.Parallel()
	userId, tearDown := setUp()
	defer tearDown()

	gdddb := wrapper.NewGddDB(db, wrapper.WithLogger(logger.NewSqlLogger(log.Default())))
	u := dao.NewUserDao(gdddb)
	affected, err := u.DeleteManySoft(context.Background(), query.C().Col("id").Eq(userId))
	if err != nil {
		panic(err)
	}
	require.NotZero(t, affected)
}
