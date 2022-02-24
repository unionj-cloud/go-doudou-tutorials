package main

import (
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	service "github.com/unionj-cloud/helloworld"
    "github.com/unionj-cloud/helloworld/config"
	"github.com/unionj-cloud/helloworld/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()
    svc := service.NewHelloworld(conf)
	handler := httpsrv.NewHelloworldHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
