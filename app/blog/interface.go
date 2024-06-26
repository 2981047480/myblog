package blog

import (
	"context"
	"vblog/common"
)

type Service interface {
	// 增
	CreateBlog(ctx context.Context, in *CreateBlogRequest) (*Blog, error)

	// 删
	DeleteBlog(ctx context.Context, in *DeleteBlogRequest) (*Blog, error)

	// 改
	UpdateBlog(ctx context.Context, in *UpdateBlogRequest) (*Blog, error)
	UpdateBlogStatus(ctx context.Context, in *UpdateBlogStatusRequest) (*Blog, error)

	// 查
	QueryBlog(ctx context.Context, in *QueryBlogRequest) (*BlogSet, error)
	DescribeBlog(ctx context.Context, in *DescribeBlogRequest) (*Blog, error)
}

func NewDeleteBlogRequest(blog_id int) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		BlogId: blog_id,
	}
}

type DeleteBlogRequest struct {
	BlogId int `json:"blog_id"`
}

func NewUpdateBlogRequest(blog_id int, update_mode common.UPDATE_MODE) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            blog_id,
		UpdateMode:        update_mode,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type UpdateBlogRequest struct {
	BlogId     int `json:"blog_id" validate:"required"`
	UpdateMode common.UPDATE_MODE

	// 修改完了需要改一下文章的状态和时间属性
	*CreateBlogRequest `validate:"required"`
}

func (r *UpdateBlogRequest) Validate() error {
	return common.Validate(r)
}

func NewUpdateBlogStatusRequest(blog_id, status int) *UpdateBlogStatusRequest {
	return &UpdateBlogStatusRequest{
		BlogId:                  blog_id,
		ChangeBlogStatusRequest: NewChangeBlogStatusRequest(status),
	}
}

type UpdateBlogStatusRequest struct {
	BlogId int `json:"blog_id"`

	// 主要是改文章的状态，所以需要id和下面的对象
	*ChangeBlogStatusRequest
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageRequest: common.NewPageRequest(),
	}
}

type QueryBlogRequest struct {
	*common.PageRequest

	// 要查询的关键字
	Keywords string
}

func NewDescribeBlogRequest(blog_id int) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		BlogId: blog_id,
	}
}

type DescribeBlogRequest struct {
	BlogId int `json:"blog_id"`
}
