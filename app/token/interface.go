package token

import (
	"context"
)

type Service interface {
	IssueToken(ctx *context.Context, in *IssueTokenRequest) (*Token, error)
	// token的颁发
	RevolkToken(ctx *context.Context, in *RevolkTokenRequest) (*Token, error)
	// token的取消

	ValicateToken(ctx *context.Context, in *ValicateTokenRequest) (*Token, error)
	// token的验证
}

func NewIssueTokenRequest(username string, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
	}
}

type IssueTokenRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsRemember bool   `json:"is_remember"`
}

func NewRevolkTokenRequest(access_token string, refresh_token string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		Access_token:  access_token,
		Refresh_token: refresh_token,
	}
}

type RevolkTokenRequest struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

func NewValicateTokenRequest(access_token string) *ValicateTokenRequest {
	return &ValicateTokenRequest{
		Access_token: access_token,
	}
}

type ValicateTokenRequest struct {
	Access_token string `json:"access_token"`
	// Refresh_token string `json:"refresh_token"` 这里一开始想的有点问题 应该不用refresh token 就像做爬虫的时候只需要cookeis就可以了
}
