package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// 这里引入mysql数据库配置 后面svc.ServiceContext中会用到
	// 这里的结构和etc配置文件中保持一致
	Mysql struct {
		DataSource string
	}
	Cache cache.CacheConf
}
