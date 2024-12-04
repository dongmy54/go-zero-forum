package svc

import (
	"forum/service/comment/cmd/rpc/comment"
	"forum/service/mq/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CommentRpc comment.Comment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		CommentRpc: comment.NewComment(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}
