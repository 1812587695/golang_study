package app

import (
	"github.com/gin-gonic/gin"

	"hytx_manager/pkg/e"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}

func (g *Gin) ResponseMsg(httpCode int, errMessage string, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  errMessage,
		"data": data,
	})

	return
}

func (g *Gin) ResponseMsgStatus(httpCode int, errMessage string, data interface{}, status bool) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  errMessage,
		"data": data,
		"status": status,
	})

	return
}

func (g *Gin) Success(data interface{}){
	g.C.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"msg" : "success",
		"data" : data,
	})
}