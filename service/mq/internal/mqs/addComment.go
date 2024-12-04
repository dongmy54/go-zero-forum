package kq

import (
	"context"
	"forum/service/mq/internal/svc"
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

	// var message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage
	// if err := json.Unmarshal([]byte(val), &message); err != nil {
	// 	logx.WithContext(l.ctx).Error("AddCommentMq->Consume Unmarshal err : %v , val : %s", err, val)
	// 	return err
	// }

	// if err := l.execService(message); err != nil {
	// 	logx.WithContext(l.ctx).Error("AddCommentMq->execService  err : %v , val : %s , message:%+v", err, val, message)
	// 	return err
	// }

	return nil
}
