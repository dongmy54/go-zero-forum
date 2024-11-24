package errorx

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

// 错误码响应
type CodeErrorResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
		Detail:  e.Detail,
	}
}

func ErrorHandle() func(err error) (int, interface{}) {
	return func(err error) (int, interface{}) {
		//错误返回
		errcode := uint32(50000) // 默认的code
		errmsg := "服务器开小差啦，稍后再来试一试"
		DetailMsg := "" // 详细信息默认为空

		causeErr := errors.Cause(err)           // err类型
		if e, ok := causeErr.(*CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.Code
			errmsg = e.Message
			DetailMsg = e.Detail
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					msgs := strings.SplitN(gstatus.Message(), "|||", 2)
					errmsg = msgs[0]
					if len(msgs) == 2 {
						DetailMsg = msgs[1]
					}
				}
			}
		}

		return http.StatusBadRequest, &CodeErrorResponse{
			Code:    errcode,
			Message: errmsg,
			Detail:  DetailMsg,
		}
	}
}
