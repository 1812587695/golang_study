package controller

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hytx_manager/models"
	"hytx_manager/pkg/setting"
	"strconv"
	"time"
)

func DistributorList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))
	data, paginate := models.DistributorList(page, limit, c)
	render(c, gin.H{
		"paginate": paginate,
		"data":     data,
	})
}

func DistributorGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	res := models.DistributorGet(id)
	render(c, res)
}

func DistributorAdd(c *gin.Context) {
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	code := c.PostForm("code")
	idCard := c.PostForm("id_card")
	idCardImgFront := c.PostForm("id_card_img_front")
	idCardImgBack := c.PostForm("id_card_img_back")
	region := com.StrTo(c.PostForm("region")).MustInt()
	pwd := c.PostForm("password")
	bankName := c.PostForm("bank_name")
	bankCardNo := c.PostForm("bank_card_no")
	bankUserName := c.PostForm("bank_user_name")

	maibao := com.StrTo(c.PostForm("maibao")).MustFloat64()

	if maibao < 0{
		fail(c, "脉宝数量错误")
		return
	}

	var cp models.Distributor
	models.DB.Model(&models.Distributor{}).Where("phone=?", phone).First(&cp)
	if phone == cp.Phone {
		fail(c, "手机号已被使用")
		return
	}
	if !models.NewSMSValidate(phone).ValidateSMS(code) {
		fail(c, "验证码错误或超时")
		return
	}

	var cp2 models.Distributor
	models.DB.Model(&models.Distributor{}).Where("region_id=?", region).First(&cp2)
	if cp2.RegionId == region {
		fail(c, "该地区已经有经销商了")
		return
	}

	var city struct{
		Name string
		Pid int
	}
	models.DB.Table("places").Select("name,pid").Where("id=?", region).First(&city)
	if city.Pid != 100000 {
		fail(c, "只有省会才能添加")
		return
	}
	pwd1, _ := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	data := models.Distributor{
		Maibao: maibao,
		MaibaoAll: maibao,
		RegionId:region,
		Region:city.Name,
		Phone:phone,
		Name:name,
		IdCard:idCard,
		IdCardImgFront:idCardImgFront,
		IdCardImgBack:idCardImgBack,
		Password:string(pwd1),
		BankName:bankName,
		BankUserName:bankUserName,
		BankCardNo:bankCardNo,
	}
	err := models.DB.Create(&data).Error
	if err != nil {
		fail(c, err.Error())
	}
	success(c)
}

func DistributorEdit(c *gin.Context) {
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	code := c.PostForm("code")
	idCard := c.PostForm("id_card")
	idCardImgFront := c.PostForm("id_card_img_front")
	idCardImgBack := c.PostForm("id_card_img_back")
	//region := com.StrTo(c.PostForm("region")).MustInt()
	//pwd := c.PostForm("password")
	bankName := c.PostForm("bank_name")
	bankCardNo := c.PostForm("bank_card_no")
	bankUserName := c.PostForm("bank_user_name")

	id := com.StrTo(c.PostForm("id")).MustInt()

	var e models.Distributor
	models.DB.Model(&models.Distributor{}).Where("id=?", id).First(&e)

	if phone != "" && phone != e.Phone{
		//更换手机号了
		var cp models.Distributor
		models.DB.Model(&models.Distributor{}).Where("phone=?", phone).First(&cp)
		if phone != cp.Phone {
			fail(c, "手机号已被使用")
			return
		}
		if !models.NewSMSValidate(phone).ValidateSMS(code) {
			fail(c, "验证码错误或超时")
			return
		}
		e.Phone = phone
	}

	e.Name = name
	e.IdCard = idCard
	e.IdCardImgFront = idCardImgFront
	e.IdCardImgBack = idCardImgBack
	e.BankName = bankName
	e.BankCardNo = bankCardNo
	e.BankUserName = bankUserName

	err := models.DB.Model(&models.Distributor{}).Where("id=?", id).Update(e).Error
	if err != nil {
		fail(c, err.Error())
	}
	success(c)
}

func DistributorDisable(c *gin.Context) {
	status := com.StrTo(c.PostForm("status")).MustInt()
	id := com.StrTo(c.PostForm("id")).MustInt()
	err := models.DB.Model(&models.Distributor{}).Where("id=?", id).Update("status", status).Error
	if err != nil {
		fail(c, err.Error())
	}
	success(c)
}

func DIstributorLog(c *gin.Context) {
	id := c.Query("id")
	phone  := c.Query("phone")
	var res []*struct{
		Phone string `json:"phone"`
		Name string `json:"name"`
		Type int `json:"type"`
		Amount float64 `json:"amount"`
		CreatedAt time.Time `json:"created_at"`
	}
	db := models.DB.Table("maibao_distributor_log").Where("distributor_id=?", id)
	if phone != "" {
		db = db.Where("phone=?", phone)
	}
	db.Scan(&res)
	render(c, res)
}

func DistributorProfitLog(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))

	data, paginate := models.DistributorProfitLog(page, limit, c)
	render(c, gin.H{
		"paginate": paginate,
		"data":     data,
	})
}

func DistributorProfitLogList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))

	id := com.StrTo(c.Query("id")).MustInt()

	data, paginate := models.DistributorProfitLogList(page, limit, c, id)
	render(c, gin.H{
		"paginate": paginate,
		"data":     data,
	})
}