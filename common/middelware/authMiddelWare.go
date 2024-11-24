package middelware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

// AuthMiddleware 是一个简单的认证中间件 这里暂时不用
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Debug("===========api中间授权中间件======")
		// 这里可以添加你的认证逻辑，例如检查请求头中的token
		// token := r.Header.Get("Authorization")
		// if token != "expected-token" {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }

		// 认证通过，调用下一个处理器
		next(w, r)
	}
}
