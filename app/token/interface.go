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

func IssueToken(ctx *context.Context, in *IssueTokenRequest) (*Token, error) {
	return nil, nil
}

type IssueTokenRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsRemember bool   `json:"is_remember"`
}

func RevolkToken(ctx *context.Context, in *IssueTokenRequest) (*Token, error) {
	return nil, nil
}

type RevolkTokenRequest struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

func ValicateToken(ctx *context.Context, in *IssueTokenRequest) (*Token, error) {
	return nil, nil
}

type ValicateTokenRequest struct {
	Access_token string `json:"access_token"`
	// Refresh_token string `json:"refresh_token"` 这里一开始想的有点问题 应该不用refresh token 就像做爬虫的时候只需要cookeis就可以了
}
