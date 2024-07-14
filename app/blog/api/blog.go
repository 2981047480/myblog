package api

import (
	"strconv"
	"vblog/app/blog"
	"vblog/common"
	"vblog/response"

	"github.com/gin-gonic/gin"
)

func (h *BlogApiHandler) Register(appRouter gin.IRouter) {
	appRouter.POST("/", h.CreateBlog)
	appRouter.GET("/", h.QueryBlog)
	appRouter.PATCH("/:id", h.PatchUpdateBlog)
	appRouter.GET("/:id", h.DescribeBlog)
	appRouter.POST("/:id", h.PostUpdateBlog)
	appRouter.DELETE("/:id", h.DeleteBlog)
}

func (h *BlogApiHandler) CreateBlog(ctx *gin.Context) {
	// 1、获取用户请求
	req := blog.NewCreateBlogRequest()
	if err := ctx.Bind(req); err != nil {
		response_failed(ctx, err)
	}
	// req.Author = ctx.Query("author")
	// req.Content = ctx.Query("content")
	// req.CreateBy = ctx.Query("author")
	// req.Summary = ctx.Query("summary")
	// req.Tags = ctx.GetStringMapString("tags")
	// req.Title = ctx.Query("title")

	// 2、业务处理
	c := ctx.Request.Context()
	b, err := h.svc.CreateBlog(c, req)
	// if err != nil {
	// 	response.Failed(err, ctx)
	// }
	response_failed(ctx, err)

	// 3、返回结果
	response.Success(b, ctx)
}

func (h *BlogApiHandler) PostUpdateBlog(ctx *gin.Context) {
	// 1、获取用户请求
	// var err error
	blog_id, err := strconv.Atoi(ctx.Param("id"))
	response_failed(ctx, err)
	update_mode := common.UPDATE_MODE_POST
	req := blog.NewUpdateBlogRequest(blog_id, update_mode)
	err = ctx.Bind(req.CreateBlogRequest)
	response_failed(ctx, err)

	// 2、业务处理
	// c := ctx.Request.Context()
	b, err := h.svc.UpdateBlog(ctx, req)
	response_failed(ctx, err)

	// 3、返回结果
	response.Success(b, ctx)
}

func (h *BlogApiHandler) PatchUpdateBlog(ctx *gin.Context) {
	// 1、获取用户请求
	blog_id, err := strconv.Atoi(ctx.Param("id"))
	response_failed(ctx, err)
	req := blog.NewUpdateBlogRequest(blog_id, common.UPDATE_MODE_PATCH)
	err = ctx.Bind(req.CreateBlogRequest)
	response_failed(ctx, err)

	// 2、业务处理
	b, err := h.svc.UpdateBlog(ctx, req)
	response_failed(ctx, err)

	// 3、返回结果
	response.Success(b, ctx)
}

func (h *BlogApiHandler) QueryBlog(ctx *gin.Context) {
	// 1、获取用户请求
	req := blog.NewQueryBlogRequest()
	req.PageRequest = common.NewPageRequestFromGin(ctx)
	req.Keywords = ctx.Query("keywords")

	// 2、业务处理
	set, err := h.svc.QueryBlog(ctx, req)
	response_failed(ctx, err)

	// 3、返回结果
	response.Success(set, ctx)
}

func (h *BlogApiHandler) DescribeBlog(ctx *gin.Context) {
	// 1、获取用户请求
	blog_id_str := ctx.Query("id")
	var blog_id int
	if blog_id_str != "" {
		blog_id, _ = strconv.Atoi(blog_id_str)
	}
	req := blog.NewDescribeBlogRequest(blog_id)

	// 2、业务处理
	set, err := h.svc.DescribeBlog(ctx, req)
	response_failed(ctx, err)

	// 3、返回结果
	response.Success(set, ctx)
}

func (h *BlogApiHandler) UpdateBlogStatus(ctx *gin.Context) {
	// 1、获取用户请求
	blog_id_str := ctx.Param("id")
	var blog_id int
	status_str := ctx.Query("status")
	var status int
	if blog_id_str != "" {
		blog_id, _ = strconv.Atoi(blog_id_str)
	}
	if status_str != "" {
		status, _ = strconv.Atoi(status_str)
	}
	req := blog.NewUpdateBlogStatusRequest(blog_id, status)
	if err := ctx.Bind(req); err != nil {
		response_failed(ctx, err)
	}

	// 2、业务处理
	b, err := h.svc.UpdateBlogStatus(ctx, req)
	response_failed(ctx, err)

	// 3、返回结果
	response.Success(b, ctx)
}

func (h *BlogApiHandler) DeleteBlog(ctx *gin.Context) {
	// 1、获取用户请求
	blog_id_str := ctx.Param("id")
	var blog_id int
	blog_id, _ = strconv.Atoi(blog_id_str)
	req := blog.NewDeleteBlogRequest(blog_id)

	// 2、业务处理
	b, err := h.svc.DeleteBlog(ctx, req)
	response_failed(ctx, err)

	// 3、返回结果
	response.Success(b, ctx)
}

func response_failed(ctx *gin.Context, err error) {
	if err != nil {
		response.Failed(err, ctx)
	}
}
