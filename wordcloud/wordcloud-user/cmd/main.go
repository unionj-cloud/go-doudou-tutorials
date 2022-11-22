package main

import (
	"github.com/sirupsen/logrus"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/db"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/internal/middleware"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/transport/httpsrv"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/etcd"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/wrapper"
)

func main() {
	defer etcd.CloseEtcdClient()
	conf := config.LoadFromEnv()
	conn, err := db.NewDb(conf.DbConf)
	if err != nil {
		panic(err)
	}
	defer func() {
		if conn == nil {
			return
		}
		if err := conn.Close(); err == nil {
			logrus.Infoln("Database connection is closed")
		} else {
			logrus.Warnln("Failed to close database connection")
		}
	}()

	db := wrapper.NewGddDB(conn)
	svc := service.NewUsersvc(conf, db)
	handler := httpsrv.NewUsersvcHandler(svc)
	srv := rest.NewRestServer()
	srv.AddMiddleware(middleware.Jwt(db))
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
