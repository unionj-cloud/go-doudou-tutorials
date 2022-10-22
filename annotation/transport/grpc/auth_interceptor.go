package grpc

import (
	"annotation/vo"
	"context"
	"encoding/base64"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sliceutils"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthInterceptor is a server interceptor for authentication and authorization
type AuthInterceptor struct {
	userStore vo.UserStore
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(userStore vo.UserStore) *AuthInterceptor {
	return &AuthInterceptor{
		userStore: userStore,
	}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// Stream returns a server interceptor function to authenticate and authorize stream RPC
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func parseToken(token string) (username, password string, ok bool) {
	c, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", "", false
	}
	cs := string(c)
	username, password, ok = strings.Cut(cs, ":")
	if !ok {
		return "", "", false
	}
	return username, password, true
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, fullMethod string) error {
	method := fullMethod[strings.LastIndex(fullMethod, "/")+1:]
	if !MethodAnnotationStore.HasAnnotation(method, "@role") {
		return nil
	}
	token, err := grpc_auth.AuthFromMD(ctx, "Basic")
	if err != nil {
		return err
	}
	user, pass, ok := parseToken(token)
	if !ok {
		return status.Error(codes.Unauthenticated, "Provide user name and password")
	}
	role, exists := interceptor.userStore[vo.Auth{user, pass}]
	if !exists {
		return status.Error(codes.Unauthenticated, "Provide user name and password")
	}
	params := MethodAnnotationStore.GetParams(method, "@role")
	if !sliceutils.StringContains(params, role.StringGetter()) {
		return status.Error(codes.PermissionDenied, "Access denied")
	}
	return nil
}
