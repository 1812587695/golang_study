package controller

import (
	"github.com/gin-gonic/gin"
	"hytx_manager/models"
)

//公用接口
//发送短信验证码
func SendSMSCode(c *gin.Context) {
	phone, phoneExist := c.GetPostForm("phone")
	if !phoneExist || phone == "" {
		fail(c, "电话错误")
	}

	sms := models.NewSMSValidate(phone)
	if sms.ID != 0 {
		if sms.IsFrequently() {
			fail(c, "操作太频繁")
			return
		}
	}

	if !sms.SendSMS(2) {
		serverError(c)
	}
	success(c)
}
//获取城市
func GetCity (c *gin.Context) {
	pId := c.DefaultQuery("pid", "100000")
	var res []*struct{
		ID int `json:"id"`
		Name string `json:"name"`
	}
	models.DB.Table("places").Select("id,name").Where("pid=?",pId).Scan(&res)
	render(c, res)
}