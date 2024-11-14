package logic

import (
	"context"

	"forum/service/comment/cmd/api/internal/svc"
	"forum/service/comment/cmd/api/internal/types"
	"forum/service/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评论
func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CreateCommentResp, err error) {
	// 先上传一个context
	myctx := context.WithValue(l.ctx, "UserId", 12)
	myctx = context.WithValue(myctx, "UserRole", "admin")

	res, err := l.svcCtx.CommentRpc.CreateComment(myctx, &pb.CreateCommentReq{
		UserId: req.UserId,
		Desc:   req.Desc,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateCommentResp{
		Id: res.Id,
	}, nil
}
