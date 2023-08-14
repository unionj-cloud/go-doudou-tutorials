/**
* Generated by go-doudou v2.1.5.
* You can edit it as your need.
 */
package plugin

import (
	service "godoudouwork1/componentb"
	"godoudouwork1/componentb/config"
	pb "godoudouwork1/componentb/transport/grpc"
	"godoudouwork1/componentb/transport/httpsrv"
	"os"

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
	"google.golang.org/grpc"
)

var _ plugin.ServicePlugin = (*ComponentbPlugin)(nil)

type ComponentbPlugin struct {
	grpcConns []*grpc.ClientConn
}

func (receiver *ComponentbPlugin) Close() {
	for _, item := range receiver.grpcConns {
		item.Close()
	}
}

func (receiver *ComponentbPlugin) GoDoudouServicePlugin() {

}

func (receiver *ComponentbPlugin) GetName() string {
	name := os.Getenv("GDD_SERVICE_NAME")
	if stringutils.IsEmpty(name) {
		name = "cloud.unionj.Componentb"
	}
	return name
}

func (receiver *ComponentbPlugin) Initialize(restServer *rest.RestServer, grpcServer *grpcx.GrpcServer, dialCtx pipeconn.DialContextFunc) {
	conf := config.LoadFromEnv()
	svc := service.NewComponentb(conf)
	restServer.AddRoute(httpsrv.Routes(httpsrv.NewComponentbHandler(svc))...)
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
	pb.RegisterComponentbServiceServer(grpcServer, svc)
}

func init() {
	zlogger.Info().Msg("from componentb")
	plugin.RegisterServicePlugin(&ComponentbPlugin{})
}
