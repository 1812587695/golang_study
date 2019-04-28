package controller

import (
	"hytx_manager/models"
	"hytx_manager/pkg/setting"
	"strconv"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func AddEnterAd(c *gin.Context) {
	title := c.PostForm("title")
	icon := c.PostForm("icon")
	broadcast := c.PostForm("broadcast")
	money, _ := com.StrTo(c.PostForm("money")).Float64()
	survivalBegin := c.PostForm("survival_begin")
	survivalEnd := c.PostForm("survival_end")

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(icon, "icon").Message("图标不能为空")
	valid.Required(broadcast, "broadcast").Message("推广方式不能为空")
	valid.Required(money, "money").Message("金额不对")
	valid.Required(survivalBegin, "survival_begin").Message("有效期开始时间不能为空")
	valid.Required(survivalEnd, "survival_end").Message("有效期结束时间不能为空")

	if valid.HasErrors() {
		fail(c, "参数格式不对")
		return
	}

	data := make(map[string]interface{})
	data["title"] = title
	data["icon"] = icon
	data["broadcast"] = broadcast
	data["money"] = money
	data["survivalBegin"] = survivalBegin
	data["survivalEnd"] = survivalEnd
	err := models.AddEnterAdvertisement(data)
	if err != nil {
		fail(c, err.Error())
		return
	}
	success(c)
}

func GetEnterAds(c *gin.Context) {

	status := c.DefaultQuery("status", "")
	beginTime := c.DefaultQuery("begin_time", "")
	endTime := c.DefaultQuery("end_time", "")

	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	limit := com.StrTo(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize))).MustInt()

	data, paginate := models.GetEnterAds(page, limit, status, beginTime, endTime)

	render(c, gin.H{
		"paginate": paginate,
		"data":     data,
	})

}

func GetEnterAd(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		fail(c, "参数格式不对")
		return
	}
	data := models.GetEnterAd(id)

	render(c, data)
}

func EditEnterAd(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	title := c.PostForm("title")
	icon := c.PostForm("icon")
	broadcast := c.PostForm("broadcast")
	money, _ := com.StrTo(c.PostForm("money")).Float64()
	survivalBegin := c.PostForm("survival_begin")
	survivalEnd := c.PostForm("survival_end")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(icon, "icon").Message("图标不能为空")
	valid.Required(broadcast, "broadcast").Message("推广方式不能为空")
	valid.Required(money, "money").Message("金额不对")
	valid.Required(survivalBegin, "survival_begin").Message("有效期开始时间不能为空")
	valid.Required(survivalEnd, "survival_end").Message("有效期结束时间不能为空")

	if valid.HasErrors() {
		fail(c, "参数格式不对")
		return
	}
	if !models.ExistEnterAdByID(id) {
		fail(c, "数据不存在")
		return
	}
	data := make(map[string]interface{})
	if title != "" {
		data["title"] = title
	}
	if icon != "" {
		data["icon"] = icon
	}
	if broadcast != "" {
		data["broadcast"] = broadcast
	}
	if money > 0 {
		data["money"] = money
	}
	if survivalBegin != "" {
		data["survival_begin"] = survivalBegin
	}
	if survivalEnd != "" {
		data["survival_end"] = survivalEnd
	}

	if !models.EditEnterAdByID(id, data) {
		fail(c, "数据更改失败")
		return
	}
	success(c)
}
func DeleteEnterAd(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		fail(c, "参数格式不对")
		return
	}
	if !models.ExistEnterAdByID(id) {
		fail(c, "数据不存在")
		return
	}
	if !models.DeleteEnterAd(id) {
		fail(c, "数据更改失败")
		return
	}
	success(c)
}

func DisableEnterAd(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	status := com.StrTo(c.PostForm("status")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Range(status, 1, 2, "status").Message("状态只允许1或2")
	if valid.HasErrors() {
		fail(c, "参数格式不对")
		return
	}
	if !models.ExistEnterAdByID(id) {
		fail(c, "数据不存在")
		return
	}
	data := make(map[string]interface{})
	data["status"] = status
	if !models.EditEnterAdByID(id, data) {
		fail(c, "数据更改失败")
		return
	}
	success(c)
}
