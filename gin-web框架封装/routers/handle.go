package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code" : http.StatusInternalServerError,
					"msg" : "服务器错误，请重试",
				})
				return
			}
		}()
		c.Next()
	}
}
