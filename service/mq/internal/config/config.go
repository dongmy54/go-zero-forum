package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	Redis redis.RedisConf

	// kq : pub sub
	CommentKqConf kq.KqConf

	// rpc
	CommentRpcConf zrpc.RpcClientConf
}
