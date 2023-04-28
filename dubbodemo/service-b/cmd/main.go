/**
* Generated by go-doudou v2.0.6.
* You can edit it as your need.
 */
package main

import (
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	service "service-b"
	"service-b/config"
	"service-b/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()
	svc := service.NewServiceB(conf)

	handler := httpsrv.NewServiceBHandler(svc)
	srv := rest.NewRestServer()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
