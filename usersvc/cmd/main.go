package main

import (
	"fmt"
	"github.com/gobwas/glob"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	"github.com/unionj-cloud/go-doudou/framework/tracing"
	"os"
	service "usersvc"
	"usersvc/config"
	"usersvc/db"
	"usersvc/internal/middleware"
	"usersvc/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()

	logger.Init()

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

	if os.Getenv("GDD_MODE") == "micro" {
		err := registry.NewNode()
		if err != nil {
			logrus.Panicln(fmt.Sprintf("%+v", err))
		}
		defer registry.Shutdown()
	}

	tracer, closer := tracing.Init()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	svc := service.NewUsersvc(conf, conn)

	handler := httpsrv.NewUsersvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	g := glob.MustCompile(fmt.Sprintf("{%s}", conf.BizConf.JwtIgnoreUrl))
	srv.AddMiddleware(middleware.Jwt(g))
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
