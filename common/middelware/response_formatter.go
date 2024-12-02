// middleware/response_formatter.go
package middelware

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseFormatter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Capture the response
		httpx.SetOkHandler(func(ctx context.Context, v interface{}) any {
			response := Response{
				Code: 0, // You can set your default code here
				Msg:  "success",
				Data: v,
			}
			return response
		})

		// Call the next handler
		next(w, r)
	}
}
