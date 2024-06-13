package impl_test

import (
	"context"
	"fmt"
	"testing"
	"vblog/app/token"
	"vblog/app/token/impl"

	_ "vblog/app"
)

var (
	ctx = context.Background()
)

func TestIssueToken(t *testing.T) {
	// u := impl.CreateTokenImpl{}
	// u.Init()
	u := impl.NewCreateTokenImpl()
	in := token.NewIssueTokenRequest("admin", "123456")
	token, err := u.IssueToken(&ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}
