package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	apolloconfig "github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	service "statsvc"
	"statsvc/config"
	"statsvc/transport/httpsrv"
	"strings"
)

type ConfigChangeListener struct {
}

func (c *ConfigChangeListener) OnChange(event *storage.ChangeEvent) {
	logger.Info("from OnChange")
	fmt.Println(event.Changes)
	for key, value := range event.Changes {
		fmt.Println("change key : ", key, ", value :", value)
	}
	fmt.Println(event.Namespace)
	logger.Info("from OnChange end")
}

func (c *ConfigChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//logger.Info("from OnNewestChange")
	//fmt.Println(event.Changes)
	//for key, value := range event.Changes {
	//	fmt.Println("change key : ", key, ", value :", value)
	//}
	//fmt.Println(event.Namespace)
	//logger.Info("from OnNewestChange end")
}

func main() {
	conf := config.LoadFromEnv()

	c := &apolloconfig.AppConfig{
		AppID:          "SampleApp",
		Cluster:        "default",
		IP:             "http://localhost:8080",
		NamespaceName:  "app.yml,app",
		IsBackupConfig: false,
		Secret:         "c6a6c0992a7a4390aafdccd428499e2b",
	}

	agollo.SetLogger(logger.Entry())

	client, _ := agollo.StartWithConfig(func() (*apolloconfig.AppConfig, error) {
		return c, nil
	})
	fmt.Println("初始化Apollo配置成功")

	namespaces := strings.Split(c.NamespaceName, ",")
	for _, item := range namespaces {
		cache := client.GetConfigCache(item)
		cache.Range(func(key, value interface{}) bool {
			fmt.Println("key : ", key, ", value :", value)
			return true
		})
	}

	var listener ConfigChangeListener
	client.AddChangeListener(&listener)
	defer func() {
		if client != nil {
			client.RemoveChangeListener(&listener)
		}
	}()

	svc := service.NewStatsvc(conf)
	handler := httpsrv.NewStatsvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
