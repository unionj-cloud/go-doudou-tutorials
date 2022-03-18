package main

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
	_ "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/toolkit/copier"
	. "github.com/unionj-cloud/go-doudou/toolkit/sqlext/query"
	"log"
)

type DbConfig struct {
	Host    string
	Port    string
	User    string
	Passwd  string
	Schema  string
	Charset string
}

type PageRetUser struct {
	PageRet
	Items []domain.User
}

func main() {
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
	var db *sqlx.DB
	db, err = sqlx.Connect("mysql", conn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	db.MapperFunc(strcase.ToSnake)

	u := dao.NewUserDao(db)

	avgScore, err := decimal.NewFromString("97.534")
	if err != nil {
		panic(err)
	}

	user := domain.User{
		Name:     "jack",
		Phone:    "13552053960",
		Age:      30,
		AvgScore: avgScore,
	}

	_, err = u.Insert(context.TODO(), &user)
	if err != nil {
		panic(err)
	}
	logrus.Printf("user %s's id is %d\n", user.Name, user.Id)

	got, err := u.PageMany(context.TODO(), Page{
		Orders: []Order{
			{
				Col:  "age",
				Sort: "desc",
			},
		},
		Offset: 0,
		Size:   1,
	}, C().Col("age").Gt(27))
	if err != nil {
		panic(err)
	}
	var ret PageRetUser
	err = copier.DeepCopy(got, &ret)
	if err != nil {
		panic(err)
	}
	logrus.Printf("returned user %s's id is %d\n", ret.Items[0].Name, ret.Items[0].Id)
	logrus.Printf("returned user %s's average score is %s", ret.Items[0].Name, ret.Items[0].AvgScore.String())

	_, err = u.DeleteMany(context.TODO(), C().Col("age").Gt(27))
	if err != nil {
		panic(err)
	}
}
