package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"forum/common/errorx"
	"forum/common/middelware"
	"forum/service/comment/cmd/api/internal/config"
	"forum/service/comment/cmd/api/internal/handler"
	midd "forum/service/comment/cmd/api/internal/middleware"
	"forum/service/comment/cmd/api/internal/svc"
	"forum/service/mq/queuemsg"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
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

	TestMsgSend(ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// 消息队列发送测试
func TestMsgSend(svcCtx *svc.ServiceContext) {
	go func() {
		for i := 0; i < 10; i++ {
			cmq := queuemsg.CommentMq{
				Id:      int64(i),
				Content: fmt.Sprintf("test-%d", i),
			}

			msg, err := json.Marshal(cmq)
			if err != nil {
				logx.Errorf("TestMsgSend Marshal err : %v", err)
				return
			}
			logx.Infof("TestMsgSend msg : %s", string(msg))
			svcCtx.KqueueCommentClient.Push(context.Background(), string(msg))
			time.Sleep(time.Second * 1)
		}
	}()
}
