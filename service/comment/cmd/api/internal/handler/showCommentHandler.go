package handler

import (
	"net/http"

	"forum/service/comment/cmd/api/internal/logic"
	"forum/service/comment/cmd/api/internal/svc"
	"forum/service/comment/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 展示评论
func ShowCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShowCommentLogic(r.Context(), svcCtx)
		resp, err := l.ShowComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
