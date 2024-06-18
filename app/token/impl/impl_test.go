package impl_test

import (
	"context"
	"fmt"
	"testing"
	"vblog/app/token"
	"vblog/ioc"

	_ "vblog/app"
)

var (
	ctx = context.Background()
	u   token.Service
)

func init() {
	ioc.ControllerImpl.Init()
	u = ioc.ControllerImpl.Get(token.AppName).(token.Service)
}

func TestIssueToken(t *testing.T) {
	in := token.NewIssueTokenRequest("default", "default")
	token, err := u.IssueToken(&ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}

func TestRevolkToken(t *testing.T) {
	in := token.NewRevolkTokenRequest("cpos4ma1jl97pss03vk0", "cpos4ma1jl97pss03vkg")
	token, err := u.RevolkToken(&ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}

func TestValicateToken(t *testing.T) {
	in := token.NewValicateTokenRequest("cpos4ma1jl97pss03vk0")
	token, err := u.ValicateToken(&ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}
