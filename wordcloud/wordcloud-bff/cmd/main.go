package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/transport/httpsrv"
	makerpb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/transport/grpc"
	taskpb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/transport/grpc"
	userclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/client"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/etcd"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/framework/restclient"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

func main() {
	defer etcd.CloseEtcdClient()
	conf := config.LoadFromEnv()

	tlsOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	dialOptions := []grpc.DialOption{
		tlsOption,
	}

	var userClient *userclient.UsersvcClient
	var makerClientConn *grpc.ClientConn
	var taskClientConn *grpc.ClientConn

	if os.Getenv("GDD_SERVICE_DISCOVERY_MODE") != "" {
		userProvider := etcd.NewSWRRServiceProvider("wordcloud-usersvc_rest")
		userClient = userclient.NewUsersvcClient(restclient.WithProvider(userProvider))

		makerClientConn = etcd.NewSWRRGrpcClientConn("wordcloud-makersvc_grpc", dialOptions...)
		defer makerClientConn.Close()

		taskClientConn = etcd.NewSWRRGrpcClientConn("wordcloud-tasksvc_grpc", dialOptions...)
		defer taskClientConn.Close()
	} else {
		userClient = userclient.NewUsersvcClient()

		serverAddr := os.Getenv("WORDCLOUDMAKER")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var err error
		makerClientConn, err = grpc.DialContext(ctx, serverAddr, dialOptions...)
		if err != nil {
			zlogger.Panic().Err(err).Msgf("[go-doudou] failed to connect to server %s", serverAddr)
		}
		defer makerClientConn.Close()

		serverAddr = os.Getenv("WORDCLOUDTASK")
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		taskClientConn, err = grpc.DialContext(ctx, serverAddr, dialOptions...)
		if err != nil {
			zlogger.Panic().Err(err).Msgf("[go-doudou] failed to connect to server %s", serverAddr)
		}
		defer taskClientConn.Close()
	}

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

	svc := service.NewWordcloudBff(conf, minioClient,
		makerpb.NewWordcloudMakerServiceClient(makerClientConn),
		taskpb.NewWordcloudTaskServiceClient(taskClientConn),
		userClient)
	handler := httpsrv.NewWordcloudBffHandler(svc)
	srv := rest.NewRestServer()
	srv.AddMiddleware(httpsrv.Auth(userClient))
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
