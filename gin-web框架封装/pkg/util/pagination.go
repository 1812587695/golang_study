package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"hytx_manager/pkg/setting"
	"math"
)

func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}

type Paginate struct {
	Page int `json:"page"` //页码
	Limit int `json:"limit"` //每页条目
	Count int `json:"count"` //总记录数
}

func (p *Paginate) SetPage(page int) {
	p.Page = page
}

func (p *Paginate) SetLimit(limit int) {
	p.Limit = limit
}

func (p *Paginate) Total() int {
	total := (p.Count + p.Limit - 1) / p.Limit
	total = int(math.Floor(float64(total)))
	if total == 0 {
		total = 1
	}
	return total
}

func (p *Paginate) Offset() int {
	return (p.Page - 1) * p.Limit
}

func (p *Paginate) SetCount(count int) {
	p.Count = count
}