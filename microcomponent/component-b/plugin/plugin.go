package plugin

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
	"github.com/unionj-cloud/go-doudou/v2/framework/plugin"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/pipeconn"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	service "github.com/wubin1989/microcomponent/component-b"
	"github.com/wubin1989/microcomponent/component-b/config"
	pb "github.com/wubin1989/microcomponent/component-b/transport/grpc"
	"github.com/wubin1989/microcomponent/component-b/transport/httpsrv"
	"google.golang.org/grpc"
	"os"
)

var _ plugin.ServicePlugin = (*ComponentBPlugin)(nil)

type ComponentBPlugin struct {
}

func (receiver *ComponentBPlugin) Close() {

}

func (receiver *ComponentBPlugin) GoDoudouServicePlugin() {

}

func (receiver *ComponentBPlugin) GetName() string {
	name := os.Getenv("GDD_SERVICE_NAME")
	if stringutils.IsEmpty(name) {
		name = "cloud.unionj.ComponentB"
	}
	return name
}

func (receiver *ComponentBPlugin) Initialize(restServer *rest.RestServer, grpcServer *grpcx.GrpcServer, _ pipeconn.DialContextFunc) {
	conf := config.LoadFromEnv()
	svc := service.NewComponentB(conf)
	restServer.AddRoute(httpsrv.Routes(httpsrv.NewComponentBHandler(svc))...)
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
	pb.RegisterComponentBServiceServer(grpcServer, svc)
}

func init() {
	plugin.RegisterServicePlugin(&ComponentBPlugin{})
}
