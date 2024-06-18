package exception

import (
	"vblog/common"
)

type ApiException struct {
	Code     int    `json:"code"`    // 类似50001之类的错误码
	Message  string `json:"message"` // 错误信息
	HttpCode int    `json:"-"`       // 不会出现在body里面，序列化成json，请求set进去
	common.Print
}

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

func (e *ApiException) Error() string {
	return e.Message
}

func ErrNotFound(message string) *ApiException {
	return NewApiException(404, message)
}
