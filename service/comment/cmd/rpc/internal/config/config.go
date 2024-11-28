package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 添加对应的结构体
	DB struct {
		DataSource string
	}

	Cache       cache.CacheConf
	UserRpcConf zrpc.RpcClientConf
	// 对于rpc 它的model关联放到svc层

	// 这里可以定义其它rpc的连接在这里
	// TravelRpcConf zrpc.RpcClientConf
}
