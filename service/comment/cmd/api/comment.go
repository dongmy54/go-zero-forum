package main

import (
	"flag"
	"fmt"

	"forum/common/errorx"
	"forum/common/middelware"
	"forum/service/comment/cmd/api/internal/config"
	"forum/service/comment/cmd/api/internal/handler"
	midd "forum/service/comment/cmd/api/internal/middleware"
	"forum/service/comment/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)

	// 添加api全局中间件 还可以对单个路由添加中间件
	server.Use(middelware.LoggingMiddleware)
	// server.Use(middelware.AuthMiddleware)

	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	// 做一些数据验证
	dataVaildMidd := midd.DataVaildateMiddleware(ctx)
	server.Use(dataVaildMidd)
	server.Use(middelware.ResponseFormatter)

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(errorx.ErrorHandle())

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
