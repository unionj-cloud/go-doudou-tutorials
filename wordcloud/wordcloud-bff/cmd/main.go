package main

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/slok/goresilience/metrics"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/transport/httpsrv"
	makerclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/client"
	taskclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/client"
	userclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/client"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	"github.com/unionj-cloud/go-doudou/framework/tracing"
	"os"
)

func main() {
	conf := config.LoadFromEnv()

	var userClient *userclient.UsersvcClient
	var makerClient *makerclient.WordcloudMakerClient
	var taskClient *taskclient.WordcloudTaskClient

	if os.Getenv("GDD_MODE") == "micro" {
		err := registry.NewNode()
		if err != nil {
			logrus.Panicln(fmt.Sprintf("%+v", err))
		}
		defer registry.Shutdown()
		userProvider := ddhttp.NewMemberlistServiceProvider("wordcloud-usersvc")
		userClient = userclient.NewUsersvcClient(ddhttp.WithProvider(userProvider))

		makerProvider := ddhttp.NewMemberlistServiceProvider("wordcloud-makersvc")
		makerClient = makerclient.NewWordcloudMakerClient(ddhttp.WithProvider(makerProvider))

		taskProvider := ddhttp.NewMemberlistServiceProvider("wordcloud-tasksvc")
		taskClient = taskclient.NewWordcloudTaskClient(ddhttp.WithProvider(taskProvider))
	} else {
		userClient = userclient.NewUsersvcClient()
		makerClient = makerclient.NewWordcloudMakerClient()
		taskClient = taskclient.NewWordcloudTaskClient()
	}

	tracer, closer := tracing.Init()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	rec := metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)

	userClientProxy := userclient.NewUsersvcClientProxy(userClient, rec)
	makerClientProxy := makerclient.NewWordcloudMakerClientProxy(makerClient, rec)
	taskClientProxy := taskclient.NewWordcloudTaskClientProxy(taskClient, rec)

	endpoint := conf.BizConf.OssEndpoint
	accessKeyID := conf.BizConf.OssKey
	secretAccessKey := conf.BizConf.OssSecret
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
	}

	svc := service.NewWordcloudBff(conf, minioClient, makerClientProxy, taskClientProxy)
	handler := httpsrv.NewWordcloudBffHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddMiddleware(httpsrv.Auth(userClientProxy))
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
