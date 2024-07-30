package middleware

import (
	"vblog/app/token"
	"vblog/app/user"
	"vblog/response"

	"github.com/gin-gonic/gin"
)

func RequireRole(requireRoles ...user.Role) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if v, ok := ctx.Get(token.GIN_TOKEN_KEY_NAME); ok {
			for i := range requireRoles {
				requireRole := requireRoles[i]
				if v.(*token.Token).Role == requireRole {
					ctx.Next()
					return
				}
			}
		}
		response.Failed(token.ErrPermissionDeny, ctx)
		ctx.Abort()
	}
}
