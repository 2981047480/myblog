package impl_test

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	_ "vblog/app"
	"vblog/app/blog"
	"vblog/common"
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
	blog, err := blogimpl.CreateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(blog)
}

func TestQueryBlog(t *testing.T) {
	in := blog.NewQueryBlogRequest()
	blog, err := blogimpl.QueryBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(blog.Item[0].Title)
}

func TestDescribeBlog(t *testing.T) {
	in := blog.NewDescribeBlogRequest(1)
	blog, err := blogimpl.DescribeBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(blog.Title)
}

func TestUpdateBlogPatch(t *testing.T) {
	in := blog.NewUpdateBlogRequest(1, common.UPDATE_MODE_PATCH)
	blog, err := blogimpl.UpdateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(blog.Title, blog.Author, blog.Content, blog.CreateBy, blog.CreateTime, blog.UpdateTime, blog.Status, blog.PublishAt)
}

func TestUpdateBlogPost(t *testing.T) {
	in := blog.NewUpdateBlogRequest(1, common.UPDATE_MODE_POST)
	blog, err := blogimpl.UpdateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(blog.Title, blog.Author, blog.Content, blog.CreateBy, blog.CreateTime, blog.UpdateTime, blog.Status, blog.PublishAt)
}

func TestUpdateBlogStatus(t *testing.T) {
	in := blog.NewUpdateBlogStatusRequest(1, 2)
	blog, err := blogimpl.UpdateBlogStatus(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(blog.Title, blog.Author, blog.Content, blog.CreateBy, blog.CreateTime, blog.UpdateTime, blog.Status, blog.PublishAt)
}

func TestDeleteBlog(t *testing.T) {
	in := blog.NewDeleteBlogRequest(1)
	blog, err := blogimpl.DeleteBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(blog.Title, blog.Author, blog.Content, blog.CreateBy, blog.CreateTime, blog.UpdateTime, blog.Status, blog.PublishAt)
}

func TestAtoi(t *testing.T) {
	str := "123"
	i, err := strconv.Atoi(str)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v %[1]T, %v, %[2]T", str, i)
}
