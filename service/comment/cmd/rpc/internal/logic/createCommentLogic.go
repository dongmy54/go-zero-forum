package logic

import (
	"context"

	"forum/common/errorx"
	"forum/service/comment/cmd/rpc/internal/svc"
	"forum/service/comment/cmd/rpc/pb"
	"forum/service/comment/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建评论
func (l *CreateCommentLogic) CreateComment(in *pb.CreateCommentReq) (*pb.CreateCommentResp, error) {
	// 模拟一个自定义错误
	//return nil, errorx.NewCodeError(450000, "创建评论失败")
	// 两种写法都可以，推荐下面的写法 可以将更多的信息记录在日志中
	return nil, errors.Wrapf(errorx.COMMENT_CREATE_ERROR, "userId:%d", in.UserId)

	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CommentModel.Insert(l.ctx, &model.Comment{
		UserId: in.UserId,
		Desc:   in.Desc,
	})

	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &pb.CreateCommentResp{
		Id: lastId,
	}, nil
}
