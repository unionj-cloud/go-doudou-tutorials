package grpc

import (
	"annotation/vo"
	"context"
	"encoding/base64"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/unionj-cloud/go-doudou/v2/framework/grpcx/interceptors/grpcx_auth"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sliceutils"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ grpcx_auth.Authorizer = (*AuthInterceptor)(nil)

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

func (interceptor *AuthInterceptor) Authorize(ctx context.Context, fullMethod string) (context.Context, error) {
	method := fullMethod[strings.LastIndex(fullMethod, "/")+1:]
	if !MethodAnnotationStore.HasAnnotation(method, "@role") {
		return ctx, nil
	}
	token, err := grpc_auth.AuthFromMD(ctx, "Basic")
	if err != nil {
		return ctx, err
	}
	user, pass, ok := parseToken(token)
	if !ok {
		return ctx, status.Error(codes.Unauthenticated, "Provide user name and password")
	}
	role, exists := interceptor.userStore[vo.Auth{user, pass}]
	if !exists {
		return ctx, status.Error(codes.Unauthenticated, "Provide user name and password")
	}
	params := MethodAnnotationStore.GetParams(method, "@role")
	if !sliceutils.StringContains(params, role.StringGetter()) {
		return ctx, status.Error(codes.PermissionDenied, "Access denied")
	}
	return ctx, nil
}
