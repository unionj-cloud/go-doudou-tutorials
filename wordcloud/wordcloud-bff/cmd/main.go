package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/slok/goresilience/metrics"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/transport/httpsrv"
	makerclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/client"
	taskclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/client"
	userclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/client"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/ratelimit"
	"github.com/unionj-cloud/go-doudou/framework/ratelimit/redisrate"
	"github.com/unionj-cloud/go-doudou/framework/tracing"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/etcd"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/framework/restclient"
	"os"
)

func main() {
	defer etcd.CloseEtcdClient()
	conf := config.LoadFromEnv()

	var userClient *userclient.UsersvcClient
	var makerClient *makerclient.WordcloudMakerClient
	var taskClient *taskclient.WordcloudTaskClient

	if os.Getenv("GDD_SERVICE_DISCOVERY_MODE") != "" {
		userProvider := etcd.NewSWRRServiceProvider("wordcloud-usersvc_rest")
		userClient = userclient.NewUsersvcClient(restclient.WithProvider(userProvider))

		makerProvider := etcd.NewSWRRServiceProvider("wordcloud-makersvc_rest")
		makerClient = makerclient.NewWordcloudMakerClient(ddhttp.WithProvider(makerProvider))

		taskProvider := etcd.NewSWRRServiceProvider("wordcloud-tasksvc_rest")
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

	// 给User、Maker、Task服务的客户端增加熔断器、超时、重试等弹性与容错机制和Prometheus指标采集
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

	svc := service.NewWordcloudBff(conf, minioClient, makerClientProxy, taskClientProxy, userClientProxy)
	handler := httpsrv.NewWordcloudBffHandler(svc)
	srv := rest.NewRestServer()
	srv.AddMiddleware(httpsrv.Auth(userClientProxy))

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", conf.RedisConf.Host),
	})

	fn := redisrate.LimitFn(func(ctx context.Context) ratelimit.Limit {
		return ratelimit.PerSecondBurst(conf.ConConf.RatelimitRate, conf.ConConf.RatelimitBurst)
	})

	srv.AddMiddleware(
		ddhttp.BulkHead(conf.ConConf.BulkheadWorkers, conf.ConConf.BulkheadMaxwaittime),
		httpsrv.RedisRateLimit(rdb, fn),
	)

	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
