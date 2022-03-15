package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/transport/httpsrv"
	segclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/client"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
	"github.com/unionj-cloud/go-doudou/framework/tracing"
	"os"
)

func main() {
	conf := config.LoadFromEnv()

	var segClient *segclient.WordcloudSegClient

	if os.Getenv("GDD_MODE") == "micro" {
		err := registry.NewNode()
		if err != nil {
			logrus.Panicln(fmt.Sprintf("%+v", err))
		}
		defer registry.Shutdown()
		provider := ddhttp.NewSmoothWeightedRoundRobinProvider("wordcloud-segsvc")
		segClient = segclient.NewWordcloudSegClient(ddhttp.WithProvider(provider))
	} else {
		segClient = segclient.NewWordcloudSegClient()
	}

	segClientProxy := segclient.NewWordcloudSegClientProxy(segClient)

	tracer, closer := tracing.Init()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

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

	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	svc := service.NewWordcloudMaker(conf, segClientProxy, minioClient, browser)
	handler := httpsrv.NewWordcloudMakerHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
