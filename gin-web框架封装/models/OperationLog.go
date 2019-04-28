package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type OperationLog struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	AdminId int `json:"admin_id"`
	Ip string `json:"ip"`
	Device string `json:"device"`
	Url string `json:"url"`
	Uri string `json:"uri"`
	Name string `json:"name"`
	Region string `json:"region"`
}

func (o *OperationLog) TableName () string {
	return "operation_log"
}

func OperationLogList(page int, limit int, c *gin.Context) (interface{}, *Paginator) {
	var res []*struct{
		ID int `json:"id"`
		CreatedAt time.Time
		Phone string `json:"phone"`
		Name string `json:"name"`
		Username string `json:"user_name"`
		Region string `json:"region"`
		Device string `json:"device"`
		Url string `json:"url"`
		IP string `json:"ip"`
	}
	m := DB.Table("operation_log as o").Select("o.id,o.created_at,o.name,o.region,o.ip,o.device,u.phone,u.username,o.url, o.uri").
		Joins("left join admins as u on u.id=o.admin_id")

	if t, isExist := c.GetQuery("time"); isExist && t != ""{
		m = m.Where("DATEDIFF(o.created_at,?)=0", t)
	}
	if name, isExist := c.GetQuery("name"); isExist && name != "" {
		m = m.Where("u.username=? or u.phone=?", name, name)
	}
	var count int
	m.Count(&count)
	paginate := NewPage(page, limit, count)
	m.Offset(paginate.Offset).Limit(paginate.Limit).Order("o.created_at desc").Find(&res)
	return res, &paginate
}