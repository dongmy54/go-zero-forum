package svc

import (
	"forum/service/comment/cmd/rpc/internal/config"
	"forum/service/comment/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	// TravelRpc travel.Travel
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
