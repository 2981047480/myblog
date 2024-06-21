package response

import (
	"net/http"
	"vblog/exception"

	"github.com/gin-gonic/gin"
)

func Success(data any, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

func Failed(err error, c *gin.Context) {
	httpCode := http.StatusInternalServerError
	// 这里判断是否为内部接口异常
	if v, ok := err.(*exception.ApiException); ok {
		if v.HttpCode != 0 {
			httpCode = v.HttpCode
		}
	} else {
		// 非业务异常 支持转化为内部异常
		err = exception.ErrServerInternal(err.Error())
	}

	c.JSON(httpCode, err)
	c.Abort()
}
