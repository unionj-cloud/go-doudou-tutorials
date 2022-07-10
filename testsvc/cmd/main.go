package main

import (
	"github.com/unionj-cloud/go-doudou/framework/dou"
	service "testsvc"
	"testsvc/config"
)

func main() {
	conf := config.LoadFromEnv()
	dou.Run(service.NewTestsvc(conf))
}
