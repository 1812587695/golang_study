package controller

import (
	"github.com/gin-gonic/gin"
	"hytx_manager/models"
	"hytx_manager/pkg/app"
	"net/http"
	"github.com/astaxie/beego/validation"
	"hytx_manager/pkg/e"
	"golang.org/x/crypto/bcrypt"
	"github.com/Unknwon/com"
	"hytx_manager/pkg/logging"
	"time"
	"strconv"
	"hytx_manager/pkg/setting"
)

// 获取
func GetAdminRolePermission(c *gin.Context) {
	appG := app.Gin{c}
	// 获取admin_id
	id := com.StrTo(c.Query("id")).MustInt()
	// 获取角色id
	role_id := models.GetAdminRoleId(id)
	// 获取角色对应的权限
	data := models.GetRolePermissionList(role_id)
	appG.Response(http.StatusOK, 200, data)
}

func GetAdmins(c *gin.Context) {

	account := c.Query("account")
	role_id := com.StrTo(c.Query("roles_id")).MustInt()
	begin_time := c.Query("begin_time")
	end_time := c.Query("end_time")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))

	data, paginate := models.GetAdmins(account, role_id, begin_time, end_time, page, limit)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": gin.H{
			"paginate": paginate,
			"data":     data,
		},
	})
}

func GetAdminById(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于1！")
	data, err := models.GetAdminById(id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "查找成功！", data)

}

// 添加
func AddAdmin(c *gin.Context) {
	appG := app.Gin{c}
	account := c.PostForm("account")
	password := c.PostForm("password")
	username := c.PostForm("username")
	email := c.PostForm("email")
	qq := c.PostForm("qq")
	department := c.PostForm("department")
	position := c.PostForm("position")
	is_enabled := com.StrTo(c.PostForm("is_enabled")).MustInt()
	logged_ip := c.ClientIP()
	phone := c.PostForm("phone")
	role_id := com.StrTo(c.PostForm("role_id")).MustInt()
	valid := validation.Validation{}
	valid.Required(account, "account").Message("账号不能为空！")
	valid.Required(password, "password").Message("密码不能为空！")
	valid.Required(username, "username").Message("姓名不能为空！")

	// 密码加密
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	logging.Info(pwd)
	if err != nil {
		appG.ResponseMsg(http.StatusOK, "注册失败,密码生成错误", false)
		return
	}

	data := models.Admin{
		Account:    account,
		Password:   string(pwd),
		Username:   username,
		Email:      email,
		Qq:         qq,
		Department: department,
		Position:   position,
		IsEnabled:  is_enabled,
		LoggedIp:   logged_ip,
		Phone:      phone,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := data.AddAdmin(role_id); err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "添加成功！", "")
}

// 修改

func UpdateAdmin(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()

	account := c.PostForm("account")
	//password := c.PostForm("password")
	username := c.PostForm("username")
	email := c.PostForm("email")
	qq := c.PostForm("qq")
	department := c.PostForm("department")
	position := c.PostForm("position")
	is_enabled := com.StrTo(c.PostForm("is_enabled")).MustInt()
	logged_ip := c.ClientIP()
	phone := c.PostForm("phone")
	role_id := com.StrTo(c.PostForm("role_id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于！")
	valid.Required(account, "account").Message("账号不能为空！")
	//valid.Required(password, "password").Message("密码不能为空！")
	valid.Required(username, "username").Message("姓名不能为空！")

	// 密码加密
	//pwd, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	//if err != nil {
	//	appG.ResponseMsg(http.StatusOK, "注册失败,密码生成错误", false)
	//	return
	//}

	data := models.Admin{
		ID:         id,
		Account:    account,
		//Password:   string(pwd),
		Username:   username,
		Email:      email,
		Qq:         qq,
		Department: department,
		Position:   position,
		IsEnabled:  is_enabled,
		LoggedIp:   logged_ip,
		Phone:      phone,
		UpdatedAt:  time.Now(),
	}

	if err := data.UpdateAdmin(role_id); err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "修改成功！", "")
}

// 删除

func DeletedAdmin(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()
	if err := models.DeletedAdmin(id); err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "删除成功！", "")
}
/**
	员工账号启用禁用
 */
func AdminAble(c *gin.Context) {
	id := c.PostForm("id")
	status,_ := strconv.Atoi(c.PostForm("statsu"))
	models.DB.Model(&models.Admin{}).Where("id=?",id).Update("is_enabled",status)
	success(c)
}
