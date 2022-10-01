package main

import (
	"fmt"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	service "github.com/unionj-cloud/helloworld"
	"github.com/unionj-cloud/helloworld/config"
	pb "github.com/unionj-cloud/helloworld/transport/grpc"
	"github.com/unionj-cloud/helloworld/transport/httpsrv"
	"google.golang.org/grpc"
	"net"
)

func main() {
	conf := config.LoadFromEnv()
	svc := service.NewHelloworld(conf)
	handler := httpsrv.NewHelloworldHandler(svc)
	srv := ddhttp.NewHttpRouterSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6061))
		if err != nil {
			logger.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterHelloworldRpcServer(s, svc)
		logger.Printf("grpc server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}()

	srv.Run()
}
