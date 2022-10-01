package main

import (
	service "abc2"
	"abc2/config"
	"abc2/transport/httpsrv"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	"github.com/unionj-cloud/helloworld/client"
	pb "github.com/unionj-cloud/helloworld/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf := config.LoadFromEnv()

	// Set up a connection to the server.
	grpcConn, err := grpc.Dial("localhost:6061", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("did not connect: %v", err)
	}
	defer grpcConn.Close()

	svc := service.NewEnumDemo(conf, pb.NewHelloworldRpcClient(grpcConn), client.NewHelloworldClient())
	handler := httpsrv.NewEnumDemoHandler(svc)
	srv := ddhttp.NewHttpRouterSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
