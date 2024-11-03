package logic

import (
	"context"

	"forum/service/comment/cmd/api/internal/svc"
	"forum/service/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 展示评论
func NewShowCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowCommentLogic {
	return &ShowCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowCommentLogic) ShowComment(req *types.ShowCommentReq) (resp *types.ShowCommentResp, err error) {
	// todo: add your logic here and delete this line
	return
}
