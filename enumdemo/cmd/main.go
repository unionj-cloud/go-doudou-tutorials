package main

import (
	service "abc2"
	"abc2/config"
	"abc2/transport/httpsrv"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
)

func main() {
	conf := config.LoadFromEnv()
	svc := service.NewEnumDemo(conf)
	handler := httpsrv.NewEnumDemoHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
