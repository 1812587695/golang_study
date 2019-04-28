package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func success(c *gin.Context) {
	send(c, http.StatusOK,"操作成功!", nil)
}

func fail(c *gin.Context, message string) {
	send(c, http.StatusBadRequest, message, nil)
}

func serverError (c *gin.Context) {
	send(c, http.StatusInternalServerError, "请稍后重试!", nil)
}

func send(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code" : code,
		"msg" : message,
		"data" : data,
	})
	c.Abort()
}

func render(c *gin.Context, args ...interface{}) {
	l := len(args)
	var message string
	var data interface{}
	if l == 0 {
		message = "success"
		data = nil
	}
	if l == 1 {
		message = "success"
		data = args[0]
	}

	if l > 1 {
		message = args[1].(string)
		data = args[0]
	}
	send(c, http.StatusOK, message, data)
}