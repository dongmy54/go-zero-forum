package svc

import (
	"forum/service/post/api/internal/config"
	"forum/service/post/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	PostModel model.PostModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		PostModel: model.NewPostModel(sqlx.NewMysql(c.Mysql.DataSource), c.Cache),
	}
}
