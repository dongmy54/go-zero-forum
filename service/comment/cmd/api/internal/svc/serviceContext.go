package svc

import (
	rpcclient "forum/common/middelware/rpcClient"
	"forum/service/comment/cmd/api/internal/config"
	"forum/service/comment/cmd/rpc/comment"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CommentRpc comment.Comment // 它是zprc生成的一个接口
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		CommentRpc: comment.NewComment(zrpc.MustNewClient(c.CommentRpcConf, zrpc.WithUnaryClientInterceptor(rpcclient.ClientAddMetadataInterceptor))),
	}
}
