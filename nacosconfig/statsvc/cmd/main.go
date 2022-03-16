package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/vo"
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

	if configmgr.NacosConfigClient != nil {
		err = configmgr.NacosConfigClient.ListenConfig(vo.ConfigParam{
			DataId: "statsvc.yml",
			Group:  "DEFAULT_GROUP",
			OnChange: func(namespace, group, dataId, data string) {
				fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			},
		})
		if err != nil {
			panic(err)
		}
		defer func() {
			err := configmgr.NacosConfigClient.CancelListenConfig(vo.ConfigParam{
				DataId: "statsvc.yml",
				Group:  "DEFAULT_GROUP",
			})
			if err != nil {
				panic(err)
			}
		}()
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
