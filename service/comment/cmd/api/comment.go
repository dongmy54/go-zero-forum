package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"forum/common/errorx"
	"forum/common/middelware"
	"forum/service/comment/cmd/api/internal/config"
	"forum/service/comment/cmd/api/internal/handler"
	"forum/service/comment/cmd/api/internal/svc"
	"forum/service/comment/cmd/rpc/comment"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/status"
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

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		//错误返回
		errcode := uint32(50000) // 默认的code
		errmsg := "服务器开小差啦，稍后再来试一试"
		DetailMsg := "" // 详细信息默认为空

		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.Code
			errmsg = e.Message
			DetailMsg = e.Detail
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if errorx.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					msgs := strings.SplitN(gstatus.Message(), "|||", 2)
					errmsg = msgs[0]
					if len(msgs) == 2 {
						DetailMsg = msgs[1]
					}
				}
			}
		}

		return http.StatusBadRequest, &errorx.CodeErrorResponse{
			Code:    errcode,
			Message: errmsg,
			Detail:  DetailMsg,
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
