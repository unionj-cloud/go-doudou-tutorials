package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/zh"
	pb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/transport/grpc"
	"github.com/unionj-cloud/go-doudou/v2/framework/grpcx"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/etcd"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"google.golang.org/grpc"
)

func main() {
	defer etcd.CloseEtcdClient()
	conf := config.LoadFromEnv()

	golac := zh.NewGoThulac(conf.BizConf.ModelPath, "", false, false, false, "_")
	defer golac.Free()

	svc := service.NewWordcloudSeg(conf, golac)
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
	pb.RegisterWordcloudSegServiceServer(grpcServer, svc)
	grpcServer.Run()
}
