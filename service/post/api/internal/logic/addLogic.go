package logic

import (
	"context"

	"forum/service/post/api/internal/svc"
	"forum/service/post/api/internal/types"
	"forum/service/post/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.PostCreateRequest) (resp *types.PostCreateResponse, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.PostModel.Insert(l.ctx, &model.Post{
		UserId:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	})

	if err != nil {
		return resp, err
	}

	id, _ := result.LastInsertId()
	return &types.PostCreateResponse{Id: id}, nil
}
