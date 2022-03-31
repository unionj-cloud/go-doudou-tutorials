package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/framework/configmgr"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	service "statsvc"
	"statsvc/config"
	"statsvc/internal/nacosservicej"
	"statsvc/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()
	err := registry.NewNode()
	if err != nil {
		logrus.Panic(fmt.Sprintf("%+v", err))
	}
	defer registry.Shutdown()

	if configmgr.NacosClient != nil {
		configmgr.NacosClient.AddChangeListener(configmgr.NacosConfigListenerParam{
			DataId: "statsvc-dev",
			OnChange: func(event *configmgr.NacosChangeEvent) {
				fmt.Println("group:" + event.Group + ", dataId:" + event.DataId + fmt.Sprintf(", changes: %+v\n", event.Changes))
			},
		})
	}

	svc := service.NewStatsvc(conf,
		nacosservicej.NewEcho(
			ddhttp.WithRootPath("/nacos-service-j"),
			ddhttp.WithProvider(ddhttp.NewNacosWRRServiceProvider("nacos-service-j"))),
	)
	handler := httpsrv.NewStatsvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
