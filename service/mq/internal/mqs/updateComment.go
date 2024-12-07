package mqs

import (
	"context"
	"encoding/json"
	"forum/service/mq/internal/svc"
	"forum/service/mq/queuemsg"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCommentMq(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentMq {
	return &UpdateCommentMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommentMq) Consume(ctx context.Context, _, val string) error {
	var message queuemsg.CommentMq
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("AddCommentMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	logc.Info(l.ctx, "UpdateCommentMq->Consume message : %+v", message)
	return nil
}
