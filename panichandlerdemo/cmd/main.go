/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package main

import (
	"fmt"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"net/http"
	service "testsvc"
	"testsvc/config"
	"testsvc/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()
	svc := service.NewTestsvc(conf)
	handler := httpsrv.NewTestsvcHandler(svc)
	srv := rest.NewRestServerWithOptions(rest.WithPanicHandler(func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				if e := recover(); e != nil {
					errmsg := fmt.Sprint(e)
					zlogger.Error().Msg(errmsg)
					http.Error(w, errmsg, http.StatusInternalServerError)
				}
			}()
			inner.ServeHTTP(w, req)
		})
	}))
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
