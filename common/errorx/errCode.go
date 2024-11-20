package errorx

var (
	//全局错误码
	SERVER_COMMON_ERROR           = New(100001, "server internal error")
	REQUEST_PARAM_ERROR           = New(100002, "request param error")
	TOKEN_EXPIRE_ERROR            = New(100003, "token expired")
	DB_ERROR                      = New(100004, "db error")
	DB_UPDATE_AFFECTED_ZERO_ERROR = New(100005, "db update affected zero")

	// 评论模块错误
	COMMENT_CREATE_ERROR = New(200001, "comment create error")
)
