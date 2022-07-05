package main

import (
	service "annotation"
	"annotation/config"
	"annotation/transport/httpsrv"
	"annotation/vo"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
)

func main() {
	conf := config.LoadFromEnv()
	svc := service.NewAnnotation(conf)
	handler := httpsrv.NewAnnotationHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddMiddleware(httpsrv.Auth(vo.UserStore{
		vo.Auth{
			User: "guest",
			Pass: "guest",
		}: vo.GUEST,
		vo.Auth{
			User: "user",
			Pass: "user",
		}: vo.USER,
		vo.Auth{
			User: "admin",
			Pass: "admin",
		}: vo.ADMIN,
	}))
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
