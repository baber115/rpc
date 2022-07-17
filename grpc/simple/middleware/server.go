package middleware

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)

// auth 中间件
func GrpcAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return newGrpcAuther().Auth
}

// stream 认证
func GrpcAuthStreamServerInterceptor() grpc.StreamServerInterceptor {
	return newGrpcAuther().streamAuth
}

// internal todo
type grpcAuther struct {
	log logger.Logger
}

func newGrpcAuther() *grpcAuther {
	return &grpcAuther{
		log: zap.L().Named("grpc auther"),
	}
}

// 获取client_id && client_secret
func (g *grpcAuther) getClientCredentialsFromMeta(md metadata.MD) (clientId, clientSecret string) {
	cids := md.Get(ClientHeaderKey)
	sids := md.Get(ClientSecretKey)
	if len(cids) > 0 {
		clientId = cids[0]
	}

	if len(sids) > 0 {
		clientSecret = sids[0]
	}
	return
}

// 验证client_id &&ｃlient_secret
func (g *grpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if clientId == "" || clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret error")
	}

	if !(clientId == "admin" && clientSecret == "123456") {
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret invalidate")
	}

	return nil
}

func (g *grpcAuther) Auth(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 1、从上下文中获取数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	fmt.Println("grpc header info", md)
	// 2、获取参数
	clientId, clientSecret := g.getClientCredentialsFromMeta(md)

	// 3、验证参数
	if err := g.validateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

/**
stream认证
重写此方法，grpc.StreamServerInterceptor()
**/
func (g *grpcAuther) streamAuth(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	fmt.Println("stream 认证")
	fmt.Println(srv, info)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		panic(ok)
	}

	clientId, clientSecret := g.getClientCredentialsFromMeta(md)

	if err := g.validateServiceCredential(clientId, clientSecret); err != nil {
		panic(err)
	}

	return handler(srv, ss)
}
