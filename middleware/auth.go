package middleware

import (
	"vblog/app/token"
	"vblog/ioc"
	"vblog/response"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	// 这里其实有点想不到怎么去做
	// 大概的思路：从context里面拿到access_token，然后验证token是否合法，合法则Next，不合法则Abort
	access_token, _ := ctx.Cookie(token.COOKIE_TOKEN_KEY)
	c := ctx.Request.Context()
	tk, err := ioc.ControllerImpl.Get(token.AppName).(token.Service).ValicateToken(&c, token.NewValicateTokenRequest(access_token))
	if err != nil {
		response.Failed(err, ctx)
		ctx.Abort()
	} else {
		ctx.Set(token.GIN_TOKEN_KEY_NAME, tk)
		ctx.Next()
	}
}
