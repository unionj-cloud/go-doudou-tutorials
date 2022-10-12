package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	ddgrpc "github.com/unionj-cloud/go-doudou/framework/grpc"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	logger "github.com/unionj-cloud/go-doudou/toolkit/zlogger"
	"google.golang.org/grpc"
	service "statsvc"
	"statsvc/config"
	"statsvc/internal/nacosservicej"
	pb "statsvc/transport/grpc"
	"statsvc/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()
	svc := service.NewStatsvc(conf,
		nacosservicej.NewEcho(
			ddhttp.WithRootPath("/nacos-service-j"),
			ddhttp.WithProvider(ddhttp.NewNacosWRRServiceProvider("nacos-service-j"))),
	)

	go func() {
		grpcServer := ddgrpc.NewGrpcServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_opentracing.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				tags.StreamServerInterceptor(tags.WithFieldExtractor(tags.CodeGenRequestFieldExtractor)),
				logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(logger.Logger)),
				grpc_recovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				tags.UnaryServerInterceptor(tags.WithFieldExtractor(tags.CodeGenRequestFieldExtractor)),
				logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger.Logger)),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		)
		pb.RegisterStatsvcServiceServer(grpcServer, svc)
		grpcServer.Run()
	}()

	handler := httpsrv.NewStatsvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
