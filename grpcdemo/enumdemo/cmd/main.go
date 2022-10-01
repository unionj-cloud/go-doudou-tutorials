package main

import (
	service "abc2"
	"abc2/config"
	"abc2/transport/httpsrv"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/klauspost/compress/gzhttp"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	"github.com/unionj-cloud/helloworld/client"
	pb "github.com/unionj-cloud/helloworld/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	conf := config.LoadFromEnv()

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		logger.Fatal("cannot load TLS credentials: ", err)
	}

	//transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	transportOption := grpc.WithTransportCredentials(tlsCredentials)

	// Set up a connection to the server.
	grpcConn, err := grpc.Dial("localhost:6061", transportOption)
	if err != nil {
		logger.Fatalf("did not connect: %v", err)
	}
	defer grpcConn.Close()

	svc := service.NewEnumDemo(conf, pb.NewHelloworldRpcClient(grpcConn), client.NewHelloworldClient(ddhttp.WithClient(newClient())))
	handler := httpsrv.NewEnumDemoHandler(svc)
	srv := ddhttp.NewHttpRouterSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
