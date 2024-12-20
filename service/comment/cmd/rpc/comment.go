package main

import (
	"flag"
	"fmt"
	"forum/common/middelware/rpcserver"

	"forum/service/comment/cmd/rpc/internal/config"
	"forum/service/comment/cmd/rpc/internal/server"
	"forum/service/comment/cmd/rpc/internal/svc"
	"forum/service/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterCommentServer(grpcServer, server.NewCommentServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	// 添加错误处理拦截器
	s.AddUnaryInterceptors(rpcserver.SetErrorInterceptor)
	// 元数据处理
	s.AddUnaryInterceptors(rpcserver.MetadataInterceptor)
	s.Start()
}
