package user

import (
	"context"
	"vblog/common"
)

const (
	AppName = "user"
)

type Service interface {
	CreateUser(ctx *context.Context, in *CreateUserRequest) (*User, error)
	QueryUser(ctx *context.Context, in *QueryUserRequest) *UserSet
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		PageRequest: common.NewPageRequest(),
	}
}

// 为什么没定义在model那边？
// 我的理解是这个不是需要做持久化的数据 所以他不是model层的
type QueryUserRequest struct {
	Username string
	*common.PageRequest
}
