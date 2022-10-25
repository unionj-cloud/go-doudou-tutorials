package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	service "github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server/config"
	pb "github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server/transport/grpc"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server/transport/httpsrv"
	"github.com/unionj-cloud/go-doudou/v2/framework/grpcx"
	"github.com/unionj-cloud/go-doudou/v2/framework/grpcx/interceptors/grpcx_ratelimit"
	"github.com/unionj-cloud/go-doudou/v2/framework/ratelimit"
	"github.com/unionj-cloud/go-doudou/v2/framework/ratelimit/memrate"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/etcd"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/nacos"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"io/ioutil"
	"strings"
	"time"
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

var _ grpcx_ratelimit.KeyGetter = (*RateLimitKeyGetter)(nil)

type RateLimitKeyGetter struct {
}

func (r *RateLimitKeyGetter) GetKey(ctx context.Context, _ string) string {
	var peerAddr string
	if value, ok := grpc_ctxtags.Extract(ctx).Values()["peer.address"]; ok {
		peerAddr = value.(string)
	}
	if stringutils.IsEmpty(peerAddr) {
		if value, ok := peer.FromContext(ctx); ok {
			peerAddr = value.Addr.String()
		}
	}
	return peerAddr[:strings.LastIndex(peerAddr, ":")]
}

func main() {
	defer nacos.CloseNamingClient()
	defer etcd.CloseEtcdClient()
	conf := config.LoadFromEnv()
	svc := service.NewHelloworld(conf)

	go func() {
		//tlsCredentials, err := loadTLSCredentials()
		//if err != nil {
		//	logger.Fatal().Err(err).Msg("cannot load TLS credentials")
		//}
		mstore := memrate.NewMemoryStore(func(_ context.Context, store *memrate.MemoryStore, key string) ratelimit.Limiter {
			return memrate.NewLimiter(10, 30, memrate.WithTimer(10*time.Second, func() {
				store.DeleteKey(key)
			}))
		})
		rl := grpcx_ratelimit.NewRateLimitInterceptor(grpcx_ratelimit.WithMemoryStore(mstore))
		keyGetter := &RateLimitKeyGetter{}
		grpcServer := grpcx.NewGrpcServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_opentracing.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(zlogger.Logger)),
				rl.StreamServerInterceptor(keyGetter),
				grpc_recovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(zlogger.Logger)),
				rl.UnaryServerInterceptor(keyGetter),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		)
		pb.RegisterHelloworldServiceServer(grpcServer, svc)
		grpcServer.Run()
	}()

	handler := httpsrv.NewHelloworldHandler(svc)
	srv := rest.NewRestServer()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
