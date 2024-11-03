package logic

import (
	"context"
	"net/http"

	"forum/service/comment/cmd/api/internal/svc"
	"forum/service/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadLogic) Upload() (resp *types.UploadRes, err error) {
	// todo: add your logic here and delete this line
	file, handler, err := l.r.FormFile("myFile")
	if err != nil {
		return nil, err
	}

	l.Errorf("==========文件名：%v 大小：%d\n", handler.Filename, handler.Size)
	l.Errorf("==========file: %#v============\n", file)
	return &types.UploadRes{Code: 200}, nil
}
