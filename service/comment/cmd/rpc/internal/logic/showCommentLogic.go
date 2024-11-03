package logic

import (
	"context"

	"forum/service/comment/cmd/rpc/internal/svc"
	"forum/service/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShowCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowCommentLogic {
	return &ShowCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 展示评论
func (l *ShowCommentLogic) ShowComment(in *pb.ShowCommentReq) (*pb.ShowCommentResp, error) {
	// todo: add your logic here and delete this line
	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ShowCommentResp{
		UserId: comment.UserId,
		Desc:   comment.Desc,
	}, nil
}
