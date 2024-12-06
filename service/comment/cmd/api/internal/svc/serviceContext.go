package svc

import (
	rpcclient "forum/common/middelware/rpcClient"
	"forum/service/comment/cmd/api/internal/config"
	"forum/service/comment/cmd/rpc/comment"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CommentRpc          comment.Comment // 它是zprc生成的一个接口
	KqueueCommentClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		CommentRpc:          comment.NewComment(zrpc.MustNewClient(c.CommentRpcConf, zrpc.WithUnaryClientInterceptor(rpcclient.ClientAddMetadataInterceptor))),
		KqueueCommentClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
