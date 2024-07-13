package logic

import (
	"context"

	"forum/service/user/model"
	"forum/service/user/rpc/internal/svc"
	"forum/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == nil {
		return &user.RegisterResponse{}, status.Error(400, "该手机号已注册")
	}

	// 加这个判断是为了避免其它错误导致去创建
	if err == model.ErrNotFound {
		res, err := l.svcCtx.UserModel.Insert(l.ctx,
			&model.Users{
				Name:     in.Name,
				Mobile:   in.Mobile,
				Password: in.Password,
				Gender:   in.Gender,
			})

		if err != nil {
			return &user.RegisterResponse{}, status.Error(400, err.Error())
		}

		id, err := res.LastInsertId()
		if err != nil {
			return &user.RegisterResponse{}, status.Error(400, err.Error())
		}

		return &user.RegisterResponse{
			Id:     id,
			Name:   in.Name,
			Gender: in.Gender,
			Mobile: in.Mobile,
		}, nil
	}

	// 如果不是ErrNotFound，则说明数据库查询出错,这里报错500
	return &user.RegisterResponse{}, status.Error(500, err.Error())
}
