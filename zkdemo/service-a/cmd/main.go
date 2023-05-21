/**
* Generated by go-doudou v2.0.6.
* You can edit it as your need.
 */
package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/zk"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/framework/restclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	service "service-a"
	"service-a/config"
	"service-a/transport/httpsrv"
	"service-b/client"
	pb "service-b/transport/grpc"
	"time"
)

func main() {
	conf := config.LoadFromEnv()
	provider := zk.NewRRServiceProvider(zk.ServiceConfig{
		Name:    "cloud.unionj.ServiceB_rest",
		Group:   "",
		Version: "",
	})
	defer provider.Close()
	bClient := client.NewServiceBClient(restclient.WithProvider(provider))

	tlsOption := grpc.WithTransportCredentials(insecure.NewCredentials())

	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(100 * time.Millisecond)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
	}

	dialOptions := []grpc.DialOption{
		tlsOption,
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_opentracing.StreamClientInterceptor(),
			grpc_retry.StreamClientInterceptor(opts...),
		)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(),
			grpc_retry.UnaryClientInterceptor(opts...),
		)),
	}

	// Set up a connection to the server.
	grpcConn := zk.NewSWRRGrpcClientConn(zk.ServiceConfig{
		Name:    "cloud.unionj.ServiceB_grpc",
		Group:   "",
		Version: "",
	}, dialOptions...)
	defer grpcConn.Close()

	svc := service.NewServiceA(conf, bClient, pb.NewServiceBServiceClient(grpcConn))
	handler := httpsrv.NewServiceAHandler(svc)
	srv := rest.NewRestServer()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}