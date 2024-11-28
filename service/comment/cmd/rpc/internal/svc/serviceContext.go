package svc

import (
	rpcclient "forum/common/middelware/rpcClient"
	"forum/service/comment/cmd/rpc/internal/config"
	"forum/service/comment/model"
	"forum/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	// TravelRpc travel.Travel
	CommentModel model.CommentModel
	UserRpc      userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserRpc:      userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(rpcclient.ClientAddMetadataInterceptor))),
	}
}
