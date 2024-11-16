package main

import (
	"flag"
	"fmt"

	"forum/common/middelware"
	"forum/service/comment/cmd/api/internal/config"
	"forum/service/comment/cmd/api/internal/handler"
	"forum/service/comment/cmd/api/internal/svc"
	"forum/service/comment/cmd/rpc/comment"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)

	// 添加api全局中间件 还可以对单个路由添加中间件
	server.Use(middelware.LoggingMiddleware)
	server.Use(middelware.AuthMiddleware)

	// 初始化RPC客户端
	commentRpcClient := comment.NewComment(zrpc.MustNewClient(c.CommentRpcConf))
	// 注入一个验证的middleware
	validaMiddleWare := middelware.NewVaildateMiddleware(commentRpcClient)
	server.Use(validaMiddleWare)

	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
