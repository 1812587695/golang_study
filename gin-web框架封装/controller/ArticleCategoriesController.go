package controller

import (
	"github.com/gin-gonic/gin"
	"hytx_manager/pkg/app"
	"hytx_manager/models"
	"hytx_manager/pkg/e"
	"net/http"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"time"
)

func GetArticleCategories(c *gin.Context) {
	appG := app.Gin{c}
	data, err := models.GetArticleCategories()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "查找成功！", data)
}

func GetArticleCategoryById(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于！")

	data, err := models.GetArticleCategoriesById(id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "", data)
}

func AddArticleCategory(c *gin.Context) {
	appG := app.Gin{c}
	title := c.PostForm("title")
	sort := com.StrTo(c.PostForm("sort")).MustInt()
	is_enabled := com.StrTo(c.PostForm("is_enabled")).MustInt()
	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空！")
	valid.Min(sort, 1, "sort").Message("sort必须大于1！")
	data := models.ArticleCategories{
		Title:     title,
		Sort:      sort,
		IsEnabled: is_enabled,
		CreatedAt: time.Now(),
	}
	err := data.AddArticleCategories()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "添加成功！", "")

}

func UpdateArticleCategory(c *gin.Context) {
	appG := app.Gin{c}
	data := map[string]interface{}{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空！")
	valid.Min(id, 1, "id").Message("id必须大于1！")
	data["id"] = id
	if title := c.Query("title"); title != "" {
		data["title"] = title
	}
	if sort := c.Query("sort"); sort != "" {
		data["sort"] = sort
	}
	if is_enabled := c.Query("is_enabled"); is_enabled != "" {
		data["is_enabled"] = is_enabled
	}
	data["updated_at"] = time.Now()
	err := models.UpdateArticleCategories(data)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}

	appG.ResponseMsg(200, "修改成功！", "")
}

func DeletedArticleCategory(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()

	err := models.DeletedArticleCategories(id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "删除成功！", "")
}

func SuggestArticleCategoryList(c *gin.Context) {
	appG := app.Gin{c}
	//id := com.StrTo(c.Query("id")).MustInt()
	data, err := models.GetSuggestArticleCategoryName()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "查找成功！", data)
}

func SuggestArticleCategoryCount(c *gin.Context) {
	appG := app.Gin{c}
	//id := com.StrTo(c.Query("id")).MustInt()
	data, err := models.GetSuggestArticleCategoryNameCount()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "查找成功！", data)
}
