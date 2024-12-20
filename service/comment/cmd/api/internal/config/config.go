package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}

	// 配置rpc
	CommentRpcConf zrpc.RpcClientConf

	// kafka
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
