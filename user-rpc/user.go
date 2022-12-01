package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/reflection"

	"go-zero-rpc/user-rpc/internal/config"
	"go-zero-rpc/user-rpc/internal/server"
	"go-zero-rpc/user-rpc/internal/svc"
	"go-zero-rpc/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "user-rpc/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, server.NewUsercenterServer(ctx))

		reflection.Register(grpcServer)
	})
	defer s.Stop()

	s.AddUnaryInterceptors(TestServerInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func TestServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return handler(ctx, req)
}
