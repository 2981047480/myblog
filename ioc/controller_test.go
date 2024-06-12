package ioc_test

import (
	"fmt"
	"testing"
	"vblog/app/user"
	"vblog/app/user/impl"
	"vblog/ioc"
)

func TestRegistor(t *testing.T) {
	controller := ioc.NewController()
	controller.Register(user.AppName, &impl.UserServiceImpl{
		Test: "test",
	})

	in := controller.Get(user.AppName).(*impl.UserServiceImpl)
	fmt.Printf("%T, %[1]v", in)
}
