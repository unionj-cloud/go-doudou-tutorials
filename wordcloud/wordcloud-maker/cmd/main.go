package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/config"
	pb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/transport/grpc"
	segpb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/transport/grpc"
	"github.com/unionj-cloud/go-doudou/v2/framework/grpcx"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/etcd"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	defer etcd.CloseEtcdClient()
	conf := config.LoadFromEnv()

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

	tlsOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	dialOptions := []grpc.DialOption{
		tlsOption,
	}
	// Set up a connection to the server.
	grpcConn := etcd.NewSWRRGrpcClientConn("wordcloud-segsvc_grpc", dialOptions...)
	defer grpcConn.Close()

	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	svc := service.NewWordcloudMaker(conf, segpb.NewWordcloudSegServiceClient(grpcConn), minioClient, browser)

	grpcServer := grpcx.NewGrpcServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(zlogger.Logger)),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(zlogger.Logger)),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterWordcloudMakerServiceServer(grpcServer, svc)
	grpcServer.Run()
}
