package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/db"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/transport/httpsrv"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	"github.com/unionj-cloud/go-doudou/framework/tracing"
	"os"
)

func main() {
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

	svc := service.NewWordcloudTask(conf, conn)
	handler := httpsrv.NewWordcloudTaskHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
