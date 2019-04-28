package controller

import (
	"github.com/gin-gonic/gin"
	"hytx_manager/middleware/auth"
)

func Profile(c *gin.Context) {
	render(c, auth.User(c))
}