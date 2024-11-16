package middelware

import (
	"forum/service/comment/cmd/rpc/comment"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

// AuthMiddleware 是一个简单的认证中间件
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Debug("===========api中间授权中间件======")
		// 这里可以添加你的认证逻辑，例如检查请求头中的token
		token := r.Header.Get("Authorization")
		if token != "expected-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 认证通过，调用下一个处理器
		next(w, r)
	}
}

func NewVaildateMiddleware(commentRpc comment.Comment) rest.Middleware {
	// 嵌套多层
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Debug("===========验证数据的middelware======")
			res, err := commentRpc.ShowComment(r.Context(), &comment.ShowCommentReq{
				Id: 18})
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			logx.Debug("===========验证数据的调用数据为：======", res)
			// 认证通过，调用下一个处理器
			next(w, r)
		}
	}
}
