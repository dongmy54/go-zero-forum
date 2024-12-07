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
		// 他们都订阅comment-topic 但是不同的group<只要是这个topic的都会被消费掉>
		// PS: 不要使用相同的topic和group 否则只有其中一个会执行
		kq.MustNewQueue(c.CommentKqConf, kqMq.NewAddCommentMq(ctx, svcContext)),
		kq.MustNewQueue(c.UpdateCommentKqConf, kqMq.NewUpdateCommentMq(ctx, svcContext)),
	}

}
