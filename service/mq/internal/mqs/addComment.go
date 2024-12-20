package mqs

import (
	"context"
	"encoding/json"
	"forum/service/mq/internal/svc"
	"forum/service/mq/queuemsg"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
*
Listening to the payment flow status change notification message queue
*/
type AddCommentMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCommentMq(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentMq {
	return &AddCommentMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCommentMq) Consume(ctx context.Context, _, val string) error {
	var message queuemsg.CommentMq
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("AddCommentMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	logc.Info(l.ctx, "AddCommentMq->Consume message : %+v", message)

	return nil
}
