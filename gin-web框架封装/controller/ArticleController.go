package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"strconv"
	"hytx_manager/pkg/setting"
	"hytx_manager/models"
	"net/http"
	"hytx_manager/pkg/app"
	"github.com/astaxie/beego/validation"
	"hytx_manager/pkg/e"
)

func GetArticles(c *gin.Context) {

	status := c.Query("status")
	article_category_id := com.StrTo(c.Query("article_category_id")).MustInt()
	operation := c.Query("operation")
	begin_time := c.Query("begin_time")
	end_time := c.Query("end_time")
	keyword:= c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))
	data, paginate := models.GetArticles(status, article_category_id, operation, begin_time, end_time,keyword, page, limit)
	articleCategories, _:= models.GetArticleCateSelect()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": gin.H{
			"paginate": paginate,
			"data":     data,
			"articleCategories":articleCategories,
		},
	})
}

func GetArticleById(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于1！")
	data, err := models.GetArticleById(id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "", data)
}

func UpdateArticle(c *gin.Context) {
	appG := app.Gin{c}
	data := map[string]interface{}{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于1！")
	data["id"] = id
	article_category_id := com.StrTo(c.Query("article_category_id")).MustInt()
	status := com.StrTo(c.Query("status")).MustInt()
	if article_category_id != 0 {
		data["article_category_id"] = article_category_id
	}
	if status != 0 {
		data["status"] = status
	}
	err := models.UpdateArticle(data)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "修改成功！", "")

}

func UpdateArticleAdvertisementId(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()
	advertisement_id := com.StrTo(c.Query("ad_id")).MustInt()
	err := models.UpdateArticleAdvertisementId(id, advertisement_id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
	}
	appG.ResponseMsg(200, "修改成功！", "")
}
