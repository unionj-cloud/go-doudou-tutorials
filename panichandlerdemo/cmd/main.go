/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
	"net/http"
	service "testsvc"
	"testsvc/config"
	"testsvc/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()
	svc := service.NewTestsvc(conf)
	handler := httpsrv.NewTestsvcHandler(svc)
	//srv := rest.NewRestServer()
	srv := rest.NewRestServerWithOptions(rest.WithPanicHandler(func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				if e := recover(); e != nil {
					statusCode := http.StatusInternalServerError
					errCode := 1 // 1 indicates there is an error
					message := fmt.Sprintf("%v", e)
					if err, ok := e.(error); ok {
						switch {
						case errors.Is(err, context.Canceled):
							statusCode = http.StatusBadRequest
						case errors.Is(err, service.BookNotFoundException):
							statusCode = http.StatusNotFound
						case errors.Is(err, service.ConversionFailedException):
							statusCode = http.StatusBadRequest
						default:
							var bizError rest.BizError
							if errors.As(err, &bizError) {
								statusCode = bizError.StatusCode
								errCode = bizError.ErrCode
								message = bizError.Error()
							}
						}
					}
					w.WriteHeader(statusCode)
					if stringutils.IsEmpty(message) {
						message = http.StatusText(statusCode)
					}
					if _err := json.NewEncoder(w).Encode(struct {
						Code    int         `json:"code,omitempty"`
						Data    interface{} `json:"data,omitempty"`
						Message string      `json:"message,omitempty"`
					}{
						Code:    errCode,
						Data:    nil,
						Message: message,
					}); _err != nil {
						http.Error(w, _err.Error(), http.StatusInternalServerError)
						return
					}
				}
			}()
			inner.ServeHTTP(w, req)
		})
	}))
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
