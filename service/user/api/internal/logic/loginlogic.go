package logic

import (
	"context"
	"time"

	"forum/common/jwtx"
	"forum/service/user/api/internal/svc"
	"forum/service/user/api/internal/types"
	"forum/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}
	// 到期时间
	expireTime := time.Now().Add(time.Duration(l.svcCtx.Config.Auth.AccessExpire) * time.Second)
	secret := l.svcCtx.Config.Auth.AccessSecret
	// 生成签名token
	accessToken, err := jwtx.GenToken(res.Id, expireTime, secret)

	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token:   accessToken,
		Expired: expireTime.Unix(),
	}, nil
}
