package main

import (
	"context"
	"flag"
	"fmt"

	"forum/service/comment/cmd/rpc/internal/config"
	"forum/service/comment/cmd/rpc/internal/server"
	"forum/service/comment/cmd/rpc/internal/svc"
	"forum/service/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	// 添加拦截器
	s.AddUnaryInterceptors(RpcServiceInterceptor)
	s.Start()
}

// 服务端拦截器
func RpcServiceInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("===========rpc服务端拦截 拦截开始=================")
	fmt.Printf("req ======> %+v \n", req)
	fmt.Printf("info =====> %+v \n", info)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "missing metadata")
	}
	fmt.Println("========收到如下元数据========")
	// 处理元数据
	for k, v := range md {
		fmt.Printf("=========key: %#v value: %#v=====\n", k, v)
	}

	resp, err = handler(ctx, req)
	fmt.Printf("===========rpc服务端拦截 拦截结束=================")

	return resp, err
}

// ===========rpc服务端拦截 拦截开始=================req ======> UserId:124 Desc:"元旦快乐"
// info =====> &{Server:0xc0003bfbb0 FullMethod:/pb.comment/CreateComment}
// ========收到如下元数据========
// =========key: "userid" value: []string{"12"}=====
// =========key: "groupid" value: []string{"23"}=====
// =========key: "user-agent" value: []string{"grpc-go/1.65.0"}=====
// =========key: "grpc-accept-encoding" value: []string{"gzip"}=====
// =========key: "userrole" value: []string{"admin"}=====
// =========key: ":authority" value: []string{"127.0.0.1:8080"}=====
// =========key: "content-type" value: []string{"application/grpc"}=====
// 2024-11-14T22:01:14.636+08:00	 error 	failed to clear cache with keys: "cache:forum:comment:id:0", error: dial tcp [::1]:6379: connect: connection refused	caller=cache/cachenode.go:84	trace=aedc22d146d817fcda02b8c3069423fe	span=43137be32130548a
// ===========rpc服务端拦截 拦截结束=================2024-11-14T22:01:20.405+08:00	 slow 	[RE
