package errorx

import "fmt"

var message = make(map[uint32]string)

type CodeError struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s, Detail:%s", e.Code, e.Message, e.Detail)
}

// 初始化各种错误
func New(code uint32, msg string) *CodeError {
	// 注入message Map中
	message[code] = msg
	return &CodeError{Code: code, Message: msg}
}

// 对错误内容进行补充
func NewErrDetail(codeErr *CodeError, Detail string) *CodeError {
	codeErr.Detail = Detail
	return codeErr
}

// 默认错误信息 都一个code
func NewDefaultError(msg string) error {
	return New(SERVER_COMMON_ERROR.Code, msg)
}

// 判断是否是自定义错误
func IsCodeErr(code uint32) bool {
	if _, ok := message[code]; ok {
		return true
	} else {
		return false
	}
}

func MapErrMsg(code uint32) string {
	return message[code]
}
