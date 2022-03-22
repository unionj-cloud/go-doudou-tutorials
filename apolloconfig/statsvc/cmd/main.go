package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/framework/configmgr"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	service "statsvc"
	"statsvc/config"
	"statsvc/transport/httpsrv"
)

type ConfigChangeListener struct {
	configmgr.BaseApolloListener
}

func (c *ConfigChangeListener) OnChange(event *storage.ChangeEvent) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	if !c.SkippedFirstEvent {
		c.SkippedFirstEvent = true
		return
	}
	logger.Info("from OnChange")
	fmt.Println(event.Changes)
	for key, value := range event.Changes {
		fmt.Println("change key : ", key, ", value :", value)
	}
	fmt.Println(event.Namespace)
	logger.Info("from OnChange end")
}

func main() {
	conf := config.LoadFromEnv()

	err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	defer registry.Shutdown()

	var listener ConfigChangeListener

	configmgr.ApolloClient.AddChangeListener(&listener)

	svc := service.NewStatsvc(conf)
	handler := httpsrv.NewStatsvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
