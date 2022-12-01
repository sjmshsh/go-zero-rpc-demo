package svc

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-rpc/user-api/internal/config"
	"go-zero-rpc/user-rpc/usercenter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(Token2uidInterceptor))),
	}
}

func Token2uidInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.New(map[string]string{
		"username": "zhangsan",
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}
	return nil
}
