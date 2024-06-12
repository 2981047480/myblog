package token

import "vblog/exception"

var (
	ErrAuthFailed          = exception.NewApiException(50001, "用户名或密码不正确")
	ErrAccessTokenExpired  = exception.NewApiException(50002, "AccessToken过期")
	ErrRefreshTokenExpired = exception.NewApiException(50003, "RefreshToken过期")
)
