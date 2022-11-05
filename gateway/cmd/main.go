package main

import (
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
)

func main() {
	srv := rest.NewRestServer()
	srv.AddMiddleware(rest.Proxy(rest.ProxyConfig{}))
	srv.Run()
}
