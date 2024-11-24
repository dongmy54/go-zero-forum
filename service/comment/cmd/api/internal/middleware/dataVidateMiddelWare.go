package midd

import (
	"net/http"
	"unicode/utf8"

	"forum/common/errorx"
	"forum/service/comment/cmd/api/internal/svc"
	"forum/service/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 做一些全局的数据验证
// PS: 注意它是内部包在internal目录下
// 注入sctx 即可做一些数据验证了哦
func DataVaildateMiddleware(svctx *svc.ServiceContext) rest.Middleware {
	// 嵌套多层
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ret, err := svctx.CommentRpc.ShowComment(ctx, &pb.ShowCommentReq{Id: 23})
			if err != nil {
				logx.Error("====comment rpc 调用失败=======", err)
			}

			if utf8.RuneCountInString(ret.Desc) > 10 {
				httpx.WriteJson(w, 400, errorx.NewDefaultError("数据验证失败"))
				return
			} else {
				next(w, r) //继续走
			}
		}
	}
}
