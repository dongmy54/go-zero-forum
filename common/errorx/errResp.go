package errorx

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
