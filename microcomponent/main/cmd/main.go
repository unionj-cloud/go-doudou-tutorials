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
	"github.com/unionj-cloud/go-doudou/v2/framework/grpcx"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	componentAService "github.com/wubin1989/microcomponent/component-a"
	apb "github.com/wubin1989/microcomponent/component-a/transport/grpc"
	componentAHttpsrv "github.com/wubin1989/microcomponent/component-a/transport/httpsrv"
	componentBService "github.com/wubin1989/microcomponent/component-b"
	bpb "github.com/wubin1989/microcomponent/component-b/transport/grpc"
	componentBHttpsrv "github.com/wubin1989/microcomponent/component-b/transport/httpsrv"
	"github.com/wubin1989/microcomponent/config"
	"google.golang.org/grpc"
)

func main() {
	srv := rest.NewRestServer()
	conf := config.LoadFromEnv()
	svcA := componentAService.NewComponentA(&conf.AConf)
	srv.AddRoute(componentAHttpsrv.Routes(componentAHttpsrv.NewComponentAHandler(svcA))...)

	svcB := componentBService.NewComponentB(&conf.BConf)
	srv.AddRoute(componentBHttpsrv.Routes(componentBHttpsrv.NewComponentBHandler(svcB))...)

	go func() {
		grpcServer := grpcx.NewGrpcServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_opentracing.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				tags.StreamServerInterceptor(tags.WithFieldExtractor(tags.CodeGenRequestFieldExtractor)),
				logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(zlogger.Logger)),
				grpc_recovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				tags.UnaryServerInterceptor(tags.WithFieldExtractor(tags.CodeGenRequestFieldExtractor)),
				logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(zlogger.Logger)),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		)
		apb.RegisterComponentAServiceServer(grpcServer, svcA)
		bpb.RegisterComponentBServiceServer(grpcServer, svcB)
		grpcServer.Run()
	}()

	srv.Run()
}
