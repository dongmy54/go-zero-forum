package logic

import (
	"context"
	"errors"

	"forum/service/user/model"
	"forum/service/user/rpc/internal/svc"
	"forum/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 注册时 先去数据库查找下是否已经有这个人
	u, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		return &user.RegisterResponse{}, err
	}

	if u != nil {
		return &user.RegisterResponse{}, errors.New("该手机号已注册")
	}

	l.svcCtx.UserModel.Insert(l.ctx,
		&model.Users{
			Name:     in.Name,
			Mobile:   in.Mobile,
			Password: in.Password,
			Gender:   in.Gender,
		})

	return &user.RegisterResponse{}, nil
}
