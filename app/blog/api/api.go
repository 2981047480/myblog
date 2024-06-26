package api

import (
	"vblog/app/blog"
	"vblog/ioc"
)

func NewBlogApiHandler() *BlogApiHandler {
	return &BlogApiHandler{
		svc: ioc.ControllerImpl.Get(blog.AppName).(blog.Service),
	}
}

type BlogApiHandler struct {
	svc blog.Service
}
