package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	ddgrpc "github.com/unionj-cloud/go-doudou/framework/grpc"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	service "github.com/unionj-cloud/helloworld"
	"github.com/unionj-cloud/helloworld/config"
	pb "github.com/unionj-cloud/helloworld/transport/grpc"
	"github.com/unionj-cloud/helloworld/transport/httpsrv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

const (
	serverCertFile   = "../cert/server-cert.pem"
	serverKeyFile    = "../cert/server-key.pem"
	clientCACertFile = "../cert/ca-cert.pem"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := ioutil.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	conf := config.LoadFromEnv()
	svc := service.NewHelloworld(conf)

	go func() {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			logger.Fatal("cannot load TLS credentials: ", err)
		}
		grpcServer := ddgrpc.NewGrpcServer(grpc.Creds(tlsCredentials))
		pb.RegisterHelloworldRpcServer(grpcServer, svc)
		grpcServer.Run()
	}()

	handler := httpsrv.NewHelloworldHandler(svc)
	srv := ddhttp.NewHttpRouterSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
