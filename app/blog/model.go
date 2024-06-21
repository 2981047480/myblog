package blog

import (
	"time"
	"vblog/common"
)

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Item: []*Blog{},
	}
}

type BlogSet struct {
	Total int64
	Item  []*Blog
}

func NewBlog() *Blog {
	return &Blog{
		common.CreateNewMeta(),
		&CreateBlogRequest{
			Tags: map[string]string{},
		},
		&ChangeBlogStatusRequest{
			Status: STATUS_DRAFT,
		},
		&common.Print{},
	}
}

type Blog struct {
	*common.Meta
	*CreateBlogRequest
	*ChangeBlogStatusRequest

	*common.Print
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Content:  "Vblog-第一个blog",
		Title:    "Vblog-第一个blog",
		Summary:  "Vblog-第一个blog",
		Author:   "Zephyr Zhao",
		CreateBy: "Zephyr Zhao",
		Tags:     map[string]string{},
	}
}

type CreateBlogRequest struct {
	Content  string            `json:"content" gorm:"column:content" validate:"required"`
	Summary  string            `json:"summary" gorm:"column:summary"`
	Author   string            `json:"author" gorm:"column:author" validate:"required"`
	Title    string            `json:"title" gorm:"column:title" validate:"required"`
	CreateBy string            `json:"create_by" gorm:"column:create_by"`
	Tags     map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
}

func (r *CreateBlogRequest) Validate() error {
	return common.Validate(r)
}

type ChangeBlogStatusRequest struct {
	PublishAt int64 `json:"published_at" gorm:"column:published_at"`
	Status    int   `json:"status" gorm:"column:status"`
}

func (r *ChangeBlogStatusRequest) SetBlogStatus(status int) {
	r.Status = status

	switch r.Status {
	case STATUS_PUBLISHD:
		r.PublishAt = time.Now().Unix()
	}
}
