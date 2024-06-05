package impl_test

import (
	"context"
	"fmt"
	"testing"
	"vblog/app/user"
	"vblog/app/user/impl"
)

var ctx = context.Background()

func TestUserFormat(t *testing.T) {
	fmt.Println(user.CreateNewUser(user.CreateNewUserRequest()).String())
}

func TestHashPassword(t *testing.T) {
	u := user.CreateNewUserRequest()
	u.HashPassword()
	fmt.Println(u.Password)
}

func TestValidatePassword(t *testing.T) {
	u := user.CreateNewUserRequest()
	u.HashPassword()
	err := u.ValidatePassword(u.Password)
	fmt.Println(err != nil)
}

func TestCreateUser(t *testing.T) {
	u := impl.UserServiceImpl{}
	u.Init()
	in := user.CreateNewUserRequest()
	u.CreateUser(&ctx, in)
}

func TestQueryUser(t *testing.T) {
	u := impl.UserServiceImpl{}
	u.Init()
	in := user.NewQueryUserRequest()
	set, _ := u.QueryUser(&ctx, in)
	fmt.Printf("%#v\n %v", set, set.Items[0])
}
