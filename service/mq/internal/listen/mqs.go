// 相当于各个handler与各个mqs的连接文件

package listen

import (
	"context"
	"forum/service/mq/internal/config"
	kqMq "forum/service/mq/internal/mqs"
	"forum/service/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.CommentKqConf, kqMq.NewAddCommentMq(ctx, svcContext)),
		//.....
	}

}
