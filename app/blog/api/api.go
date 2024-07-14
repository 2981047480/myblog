package api

import (
	"vblog/app/blog"
	"vblog/common/config"
	"vblog/ioc"
)

func init() {
	ioc.Api.Register(blog.AppName, &BlogApiHandler{})
}

func NewBlogApiHandler() *BlogApiHandler {
	return &BlogApiHandler{
		svc: ioc.ControllerImpl.Get(blog.AppName).(blog.Service),
	}
}

type BlogApiHandler struct {
	svc blog.Service
}

func (h *BlogApiHandler) Init() error {
	h.svc = ioc.ControllerImpl.Get(blog.AppName).(blog.Service)

	// 注册路由
	subRouter := config.ReadDBConf(config.Filename).Application.GinRootRouter().Group("blogs")
	h.Register(subRouter)
	return nil
}
