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

func (e *ApiException) WithMessage(message string) *ApiException {
	e.Message = message
	return e
}

func (e *ApiException) WithHttpCode(code int) *ApiException {
	e.HttpCode = code
	return e
}
