package impl_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	_ "vblog/app"
	"vblog/app/blog"
	"vblog/ioc"
)

func init() {
	if err := ioc.ControllerImpl.Init(); err != nil {
		log.Panic(err)
	}
}

var (
	ctx      = context.Background()
	blogimpl = ioc.ControllerImpl.Get(blog.AppName).(blog.Service)
)

func TestCreateBlog(t *testing.T) {
	in := blog.NewCreateBlogRequest()
	blog, err := blogimpl.CreateBlog(&ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(blog)
}
