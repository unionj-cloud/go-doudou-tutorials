package plugin

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/unionj-cloud/go-doudou/v2/framework/grpcx"
	"github.com/unionj-cloud/go-doudou/v2/framework/plugin"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/pipe"
	"github.com/unionj-cloud/go-doudou/v2/framework/registry/zk"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/pipeconn"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	service "github.com/wubin1989/microcomponent/component-a"
	"github.com/wubin1989/microcomponent/component-a/config"
	apb "github.com/wubin1989/microcomponent/component-a/transport/grpc"
	"github.com/wubin1989/microcomponent/component-a/transport/httpsrv"
	bpb "github.com/wubin1989/microcomponent/component-b/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

var _ plugin.ServicePlugin = (*ComponentAPlugin)(nil)

type ComponentAPlugin struct {
	grpcConns []*grpc.ClientConn
}

func (receiver *ComponentAPlugin) Close() {
	for _, item := range receiver.grpcConns {
		item.Close()
	}
}

func (receiver *ComponentAPlugin) GoDoudouServicePlugin() {

}

func (receiver *ComponentAPlugin) GetName() string {
	return os.Getenv("GDD_SERVICE_NAME")
}

func (receiver *ComponentAPlugin) Initialize(restServer *rest.RestServer, grpcServer *grpcx.GrpcServer, dialCtx pipeconn.DialContextFunc) {
	conf := config.LoadFromEnv()
	var componentBClient bpb.ComponentBServiceClient
	if dialCtx != nil {
		grpcConn := pipe.NewGrpcClientConn(dialCtx)
		receiver.grpcConns = append(receiver.grpcConns, grpcConn)
		componentBClient = bpb.NewComponentBServiceClient(grpcConn)
	} else {
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
			Name:    "cloud.unionj.ComponentB_grpc",
			Group:   "group",
			Version: "v1.0.0",
		}, dialOptions...)
		receiver.grpcConns = append(receiver.grpcConns, grpcConn)
		componentBClient = bpb.NewComponentBServiceClient(grpcConn)
	}
	svc := service.NewComponentA(conf, componentBClient)
	restServer.AddRoute(httpsrv.Routes(httpsrv.NewComponentAHandler(svc))...)
	if grpcServer.Server == nil {
		grpcServer.Server = grpc.NewServer(
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
	}
	apb.RegisterComponentAServiceServer(grpcServer, svc)
}

func init() {
	plugin.RegisterServicePlugin(&ComponentAPlugin{})
}
