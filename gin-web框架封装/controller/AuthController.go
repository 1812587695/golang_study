package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	authM "hytx_manager/middleware/auth"
	"hytx_manager/models"
)

func Login(c *gin.Context) {
	auth := c.MustGet("auth").(authM.Auth)
	account := c.PostForm("account")
	password := c.PostForm("password")
	admin := models.FindAdminAccount(account)

	if admin == nil {
		fail(c, "账号或密码错误!")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		fail(c, "账号或密码错误!")
		return
	}
	if admin.IsEnabled == 1 {
		fail(c, "该账号禁止登录,请联系管理员!")
		return
	}
	token := auth.Login(c.Request, c.Writer, map[string]interface{}{"id": admin.ID}).(string)
	// 获取用户角色对应的权限菜单
	role_id := models.GetAdminRoleId(admin.ID)
	item := models.GetRolePermissionList(role_id)
	data := map[string]interface{}{}
	data["token"] = token
	data["item"] =item

	render(c, data)
}
