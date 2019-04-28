package controller

import (
	"hytx_manager/models"
	"hytx_manager/pkg/setting"
	"strconv"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func AddBanner(c *gin.Context) {
	title := c.PostForm("title")
	images := c.PostForm("images")
	sort := com.StrTo(c.DefaultPostForm("sort", "0")).MustInt()
	url := c.PostForm("url")
	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(images, "images").Message("图片不能为空")
	valid.Required(url, "url").Message("链接不能为空")

	if valid.HasErrors() {
		fail(c, "参数格式不对")
		return
	}

	data := make(map[string]interface{})
	data["title"] = title
	data["sort"] = sort
	data["images"] = images
	data["url"] = url
	err := models.AddBanner(data)
	if err != nil {
		fail(c, err.Error())
		return
	}
	success(c)
}
func GetBanners(c *gin.Context) {

	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	limit := com.StrTo(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize))).MustInt()

	data, paginate := models.GetBanners(page, limit)

	render(c, gin.H{
		"paginate": paginate,
		"data":     data,
	})

}
func GetBanner(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		fail(c, "参数格式不对")
		return
	}
	data := models.GetBanner(id)

	render(c, data)
}
func DeleteBanner(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		fail(c, "参数格式不对")
		return
	}
	if !models.ExistBannerByID(id) {
		fail(c, "数据不存在")
		return
	}
	if !models.DeleteBanner(id) {
		fail(c, "数据删除失败")
		return
	}
	success(c)
}
func EditBanner(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	title := c.PostForm("title")
	images := c.PostForm("images")
	sort := com.StrTo(c.DefaultPostForm("sort", "0")).MustInt()
	url := c.PostForm("url")
	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(images, "images").Message("图片不能为空")
	valid.Required(url, "url").Message("链接不能为空")

	if valid.HasErrors() {
		fail(c, "参数格式不对")
		return
	}
	if !models.ExistBannerByID(id) {
		fail(c, "数据不存在")
		return
	}
	data := make(map[string]interface{})
	if title != "" {
		data["title"] = title
	}
	data["sort"] = sort
	if images != "" {
		data["images"] = images
	}
	if url != "" {
		data["url"] = url
	}

	if !models.EditBannerByID(id, data) {
		fail(c, "数据更改失败")
		return
	}
	success(c)
}
