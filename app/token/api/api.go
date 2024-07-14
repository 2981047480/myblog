package api

import (
	"vblog/app/token"
	"vblog/common/config"
	"vblog/ioc"
	"vblog/response"

	"github.com/gin-gonic/gin"
)

func init() {
	ioc.Api.Register(token.AppName, &TokenApiHandler{})
}

func NewTokenApiHandler() *TokenApiHandler {
	return &TokenApiHandler{
		svc: ioc.ControllerImpl.Get(token.AppName).(token.Service),
	}
}

type TokenApiHandler struct {
	svc token.Service
}

func (h *TokenApiHandler) Init() error {
	h.svc = ioc.ControllerImpl.Get(token.AppName).(token.Service)

	subRouter := config.ReadDBConf(config.Filename).Application.GinRootRouter().Group("tokens")
	h.Register(subRouter)
	return nil
}

func (h *TokenApiHandler) Register(appRouter gin.IRouter) {
	appRouter.POST("/", h.Login)
	appRouter.DELETE("/", h.Logout)
}

func (h *TokenApiHandler) Login(c *gin.Context) {
	// 1、获取http请求
	req := token.NewIssueTokenRequest("", "")
	if err := c.BindJSON(req); err != nil {
		response.Failed(err, c)
	}

	// 2、业务逻辑处理
	ctx := c.Request.Context()
	tk, err := h.svc.IssueToken(&ctx, req)
	if err != nil {
		response.Failed(err, c)
		return
	}

	// 3、返回结果
	c.SetCookie(token.COOKIE_TOKEN_KEY,
		tk.AccessToken,
		tk.RefreshTokenExpiredAt,
		"/", config.Default().Application.Domain,
		false, true)

	response.Success(tk, c)
}

func (h *TokenApiHandler) Logout(c *gin.Context) {
	ak, err := c.Cookie(token.COOKIE_TOKEN_KEY)
	if err != nil {
		response.Failed(err, c)
		return
	}

	rt := c.GetHeader(token.REFRESH_HEAD_KEY)
	req := token.NewRevolkTokenRequest(ak, rt)
	ctx := c.Request.Context()
	tk, err := h.svc.RevolkToken(&ctx, req)
	if err != nil {
		response.Failed(err, c)
		return
	}

	response.Success(tk, c)

}
