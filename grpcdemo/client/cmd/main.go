package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/go-resty/resty/v2"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/klauspost/compress/gzhttp"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	service "github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/client"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/client/config"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/client/transport/httpsrv"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server/client"
	pb "github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server/transport/grpc"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/etcd"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/nacos"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/framework/restclient"
	"github.com/unionj-cloud/go-doudou/v2/framework/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"time"
)

func newClient() *resty.Client {
	client := resty.New()
	client.SetTimeout(1 * time.Minute)
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	client.SetTransport(gzhttp.Transport(&nethttp.Transport{
		RoundTripper: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialer.DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
			MaxConnsPerHost:       100000,
		},
	}))
	client.SetRetryCount(10)
	return client
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("../cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("../cert/client-cert.pem", "../cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	defer nacos.CloseNamingClient()
	defer etcd.CloseEtcdClient()
	conf := config.LoadFromEnv()

	tracer, closer := tracing.Init()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	//tlsCredentials, err := loadTLSCredentials()
	//if err != nil {
	//	logger.Fatal().Err(err).Msg("cannot load TLS credentials")
	//}

	tlsOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	//tlsOption := grpc.WithTransportCredentials(tlsCredentials)

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
	grpcConn := nacos.NewWRRGrpcClientConn(nacos.NacosConfig{
		ServiceName: "grpcdemo-server_grpc",
	}, dialOptions...)
	//grpcConn := etcd.NewRRGrpcClientConn("grpcdemo-server_grpc", dialOptions...)
	//grpcConn := etcd.NewSWRRGrpcClientConn("grpcdemo-server_grpc", dialOptions...)
	defer grpcConn.Close()

	//restProvider := nacos.NewWRRServiceProvider("grpcdemo-server_rest")
	restProvider := etcd.NewRRServiceProvider("grpcdemo-server_rest")
	//restProvider := etcd.NewSWRRServiceProvider("grpcdemo-server_rest")
	svc := service.NewEnumDemo(conf, pb.NewHelloworldServiceClient(grpcConn),
		client.NewHelloworldClient(restclient.WithClient(newClient()), restclient.WithProvider(restProvider)))
	handler := httpsrv.NewEnumDemoHandler(svc)
	srv := rest.NewRestServer()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
