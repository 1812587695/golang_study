package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"hytx_manager/middleware/auth"
	"hytx_manager/models"
	"hytx_manager/pkg/setting"
	"strconv"
	"time"
)

func AdList(c *gin.Context) {
	status := c.Query("status")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))

	beginTime := c.Query("begin_time")
	endTime := c.Query("end_time")

	res, paginate := models.AdList(page, limit, status, beginTime, endTime)
	articleCategories, _:= models.GetArticleCateSelect()
	articleLabels, _:=models.GetArticleLabelSelect()
	render(c, gin.H{
		"paginate": paginate,
		"data":     res,
		"articleCategories":articleCategories,
		"articleLabels":articleLabels,

	})
}

func AddAd(c *gin.Context) {
	operator := auth.User(c).ID
	valid := validation.Validation{}
	typee, _ := strconv.Atoi(c.PostForm("type"))
	valid.Required(typee, "type").Message("类型不能为空")
	name := c.PostForm("name")
	valid.Required(name, "name").Message("名称不能为空")
	images := c.PostForm("images")
	valid.Required(images, "images").Message("图片不能为空")
	describe := c.PostForm("describe")
	//valid.Required(describe, "describe").Message("描述不能为空")
	url := c.PostForm("url")

	cate, _ := strconv.Atoi(c.PostForm("cate"))
	valid.Required(cate, "cate").Message("分类不能为空")
	label, _ := strconv.Atoi(c.PostForm("label"))
	valid.Required(label, "label").Message("标签不能为空")
	survivalBegin64,_ := strconv.ParseInt(c.PostForm("survival_begin"), 10, 64)
	survivalBegin := time.Unix(survivalBegin64/1e3, 0)
	valid.Required(survivalBegin, "survival_begin").Message("开始时间不能为空")
	survivalEnd64,_ := strconv.ParseInt(c.PostForm("survival_end"), 10, 64)
	survivalEnd := time.Unix(survivalEnd64/1e3, 0)
	valid.Required(survivalEnd, "survival_end").Message("结束时间不能为空")
	data := models.Advertisement{
		Operator:      operator,
		Type:          typee,
		Name:          name,
		Images:        images,
		Describe:      describe,
		Url:           url,
		Cate:          cate,
		Label:         label,
		SurvivalBegin: survivalBegin,
		SurvivalEnd:   survivalEnd,
		CreatedAt:     time.Now(),
	}
	status, statusExist := c.GetPostForm("status")
	if  statusExist && status != "" {
		data.Status, _ = strconv.Atoi(status)
	}
	if time.Now().After(survivalBegin) && time.Now().Before(survivalEnd) && status != "" {
		data.Status = 2
	}
	if err := data.Add(); err != nil {
		serverError(c)
		return
	}
	success(c)
}

func EditAd(c *gin.Context) {
	id := c.Query("id")
	data := make(map[string]interface{})

	if t, isExist := c.GetPostForm("type"); isExist && t != "" {
		data["type"] = t
	}
	if name, isExist := c.GetPostForm("name"); isExist && name != "" {
		data["name"] = name
	}
	if image, isExist := c.GetPostForm("images"); isExist && image != "" {
		data["images"] = image
	}
	if describe, isExist := c.GetPostForm("describe"); isExist && describe != "" {
		data["describe"] = describe
	}
	if url, isExist := c.GetPostForm("url"); isExist && url != "" {
		data["url"] = url
	}
	if cate, isExist := c.GetPostForm("cate"); isExist && cate != "" {
		data["cate"] = cate
	}
	if label, isExist := c.GetPostForm("label"); isExist && label != "" {
		data["label"] = label
	}
	if survivalBegin, isExist := c.GetPostForm("survival_begin"); isExist && survivalBegin != "" {

		survivalBegin64,_ := strconv.ParseInt(c.PostForm("survival_begin"), 10, 64)
		data["survival_begin"] = time.Unix(survivalBegin64/1e3, 0)
	}
	if survivalEnd, isExist := c.GetPostForm("survival_end"); isExist && survivalEnd != "" {
		survivalEnd64,_ := strconv.ParseInt(c.PostForm("survival_end"), 10, 64)
		data["survival_end"] = time.Unix(survivalEnd64/1e3, 0)
	}

	status, statusExist := c.GetPostForm("status")
	if  statusExist && status != "" {
		data["status"], _ = strconv.Atoi(status)
	}
	if statusExist && status == "2" && data["survival_end"].(time.Time).Before(time.Now()) {
		fail(c, "结束时间不能在今天之前")
		return
	}
	if time.Now().After(data["survival_begin"].(time.Time)) && time.Now().Before(data["survival_end"].(time.Time)) && status != "" {
		data["status"] = 2
	}
	data["updated_at"] = time.Now()
	 if err := models.DB.Table("advertisement").Where("id=?", id).Updates(data).Error; err != nil{
	 	serverError(c)
	 }
	success(c)
}

func AdGet(c *gin.Context) {
	id := c.Query("id")
	data := models.GetAdById(id)

	render(c, data)
}

func DElAd(c *gin.Context) {
	id := c.Query("id")
	models.DB.Delete(models.Advertisement{}, "id=?", id)
	success(c)
}


func AdBan(c *gin.Context) {
	//下架
	id := c.PostForm("id")
	models.DB.Table("advertisement").Where("id=?", id).Updates(map[string]interface{}{"status":3,"updated_at":time.Now(),})
	success(c)
}