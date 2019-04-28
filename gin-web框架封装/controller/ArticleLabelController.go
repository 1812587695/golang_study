package controller

import (
	"github.com/gin-gonic/gin"
	"hytx_manager/models"
	"hytx_manager/pkg/app"
	"net/http"
	"hytx_manager/pkg/e"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"time"
)

func GetArticleLabel(c *gin.Context) {
	appG := app.Gin{c}
	//data := models.GetArticleLabels
	data, err := models.GetArticleLabels()

	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "查找成功！", data)
}

func GetArticleLabelById(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于！")

	data, err := models.GetArticleLabelById(id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "", data)

}

func AddArticleLabel(c *gin.Context) {
	appG := app.Gin{c}
	name := c.PostForm("name")
	sort := c.PostForm("sort")
	is_enabled := com.StrTo(c.PostForm("is_enabled")).MustInt()
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空！")
	valid.Min(sort, 1, "sort").Message("sort必须大于1！")
	data := models.ArticleLabel{
		Name:      name,
		Sort:      sort,
		IsEnabled: is_enabled,
		CreatedAt: time.Now(),
	}
	err := data.AddArticleLabel()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "添加成功！", "")
}

func UpdateArticleLabel(c *gin.Context) {
	appG := app.Gin{c}
	data := map[string]interface{}{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空！")
	valid.Min(id, 1, "id").Message("id必须大于1！")
	data["id"] = id
	name := c.Query("name")
	if name != "" {
		data["name"] = name
	}
	if sort := c.Query("sort"); sort != "" {
		data["sort"] = sort
	}
	if is_enabled := c.Query("is_enabled"); is_enabled != "" {
		data["is_enabled"] = is_enabled
	}
	data["updated_at"] = time.Now()

	err := models.UpdateArticleLabel(data)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}

	appG.ResponseMsg(200, "修改成功！", "")
}

func DeletedArticleLabel(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()

	err := models.DeletedArticleLabel(id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "删除成功！", "")
}
