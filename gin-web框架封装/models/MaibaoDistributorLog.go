package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type MaibaoDistributorLog struct {
	ID int `json:"id"`
	UserId int `json:"user_id"`
	DistributorId int `json:"distributor_id"`
	Cate int `json:"cate"`
	Type int `json:"type"`
	Amount float64 `json:"amount"`
	AmountReal float64 `json:"amount_real"`
	Content string `json:"content"`
	Phone string `json:"phone"`
	Name string `json:"name"`
	Ip string `json:"ip"`
	Region string `json:"region"`
	RegionId int `json:"region_id"`
	Additional string `json:"additional"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *MaibaoDistributorLog) TableName () string {
	return "maibao_distributor_log"
}

func MaibaoLog(page int, limit int, c *gin.Context) (interface{}, *Paginator) {
	var res []*struct{
		UserId int `json:"user_id"`
		Name string `json:"name"`
		Phone string `json:"phone"`
		UserType string `json:"user_type"`
		Additional string `json:"additional"`
		Region string `json:"region"`
		Maibao float64 `json:"maibao"`
		Money float64 `json:"money"`
		Status int `json:"status"`
	}
	selec := "select * from "
	from := "(SELECT u.phone,c.user_name as `name`,p.`name` as region,m.user_id, m.user_type, sum(m.amount) as maibao, sum(m.money) as money,'个人' as additional,u.`is_enabled` as `status` FROM `maibao` m LEFT JOIN users u on u.id=m.user_id LEFT JOIN cards c on c.user_id=m.user_id LEFT JOIN places p on p.id=c.province WHERE  m.user_type='user' GROUP BY m.user_id,m.user_type " +
		" UNION " +
		"SELECT es.phone,es.`name`,p.`name` as region,m.user_id, m.user_type, sum(m.amount) as maibao, sum(m.money) as money,e.`name` as additional,e.`status` FROM `maibao` m LEFT JOIN enterprise e on e.id=m.user_id LEFT JOIN enterprise_users as eu on eu.enterprise_id=e.id and eu.identity=1 LEFT JOIN enterprise_staff es on es.id=eu.staff_id  LEFT JOIN places p on p.id=e.province_id where m.user_type='enterprise' GROUP BY m.user_id,m.user_type)" +
		" as a "

	hasWhere := false
	if region := c.Query("region"); region != "" {
		hasWhere = appendSql(&from, hasWhere)
		from += "a.`region`=" + "'"+region+"'"
	}
	if status := c.Query("status"); status != "" {
		hasWhere = appendSql(&from, hasWhere)
		from += "a.`status`=" + "'"+status+"'"
	}
	if phone := c.Query("phone"); phone != "" {
		hasWhere = appendSql(&from, hasWhere)
		from += "a.`phone`=" + "'"+phone+"'"
	}
	if name := c.Query("name"); name != "" {
		hasWhere = appendSql(&from, hasWhere)
		from += "a.`additional`=" + "'"+name+"'"
	}
	var count struct{
		Count int
	}
	countSql := "select count(*) as count from " + from
	err := Table(MaibaoDistributorLogTable).Raw(countSql).Scan(&count).Error
	if err != nil {
		print(err)
	}
	paginate := NewPage(page, limit, count.Count)
	sql := selec + from
	sql += " limit " + strconv.Itoa(paginate.Limit) + " offset " + strconv.Itoa(paginate.Offset)
	Table(MaibaoDistributorLogTable).Raw(sql).Scan(&res)
	return res, &paginate
}

func appendSql(sql *string, hasWhere bool) bool {
	if hasWhere {
		*sql += " AND "
	}else {
		*sql += " WHERE "
	}
	return true
}

func MaibaoRechargeLog(page int, limit int, c *gin.Context, userType string) (interface{}, *Paginator) {
	var res []*struct{
		UserId int `json:"user_id"`
		Name string `json:"name"`
		Phone string `json:"phone"`
		UserType string `json:"user_type"`
		Additional string `json:"additional"`
		Region string `json:"region"`
		Maibao float64 `json:"maibao"`
		Money float64 `json:"money"`
		Status int `json:"status"`
	}
	var m *gorm.DB
	if userType == "user"{
		m = DB.Table("maibao as m").Select("u.phone,c.user_name as `name`,p.`name` as region,m.user_id, m.user_type, sum(m.amount) as maibao, sum(m.money) as money,'个人' as additional,u.`is_enabled` as `status` ").
			Joins("left join users u on u.id=m.user_id").
			Joins("left join cards c on c.user_id=m.user_id").
			Joins("LEFT JOIN places p on p.id=c.province").
			Where("m.user_type='user'").
			Group("m.user_id,m.user_type")


		if region := c.Query("region"); region != "" {
			m = m.Where("p.name=?", region)
		}
		if status := c.Query("status"); status != "" {
			m = m.Where("u.is_enabled=?", status)
		}
		if phone := c.Query("phone"); phone != "" {
			m = m.Where("u.phone=?", phone)
		}
		if name := c.Query("name"); name != "" {
			m = m.Where("c.user_name=?", name)
		}
	}else {
		m = DB.Table("maibao as m").
			Select("es.phone,es.`name`,p.`name` as region,m.user_id, m.user_type, sum(m.amount) as maibao, sum(m.money) as money,e.`name` as additional,eu.`status` ").
			Joins("LEFT JOIN enterprise_staff es on es.id=m.user_id").
			Joins("LEFT JOIN enterprise e on e.id=es.enterprise_id").
			Joins("LEFT JOIN enterprise_users as eu on eu.staff_id=m.user_id").
			Joins("LEFT JOIN places p on p.id=e.province_id").
			Where("m.user_type='enterprise'").
			Group("m.user_id,m.user_type")

		if region := c.Query("region"); region != "" {
			m = m.Where("p.name=?", region)
		}
		if status := c.Query("status"); status != "" {
			m = m.Where("eu.status=?", status)
		}
		if phone := c.Query("phone"); phone != "" {
			m = m.Where("es.phone=?", phone)
		}
		if name := c.Query("name"); name != "" {
			m = m.Where("es.name=?", name)
		}
	}
	var count int
	m.Count(&count)
	paginate := NewPage(page, limit, count)
	m.Scan(&res)
	return &res, &paginate
}

func MaibaoRechargeLogByUser(page int, limit int, c *gin.Context, userType string, userId int) (interface{}, *Paginator) {
	var res []*struct{
		Name string `json:"name"`
		Phone string `json:"phone"`
		UserType string `json:"user_type"`
		Region string `json:"region"`
		Money float64 `json:"money"`
		Maibao float64 `json:"maibao"`
		Additional string `json:"additional"`
		Type string `json:"type"`
		CreatedAt time.Time `json:"created_at"`
	}
	var m *gorm.DB
	if userType == "user"{
		m = DB.Table("maibao as m").
			Select("m.`created_at`,m.`type`,m.`amount` as `maibao`,m.`money`,u.`phone`,p.`name` as `region`, '个人' as `additional`").
			Joins("left join users u on u.id=m.user_id").
			Joins("left join cards c on c.user_id=m.user_id").
			Joins("LEFT JOIN places p on p.id=c.province").
			Where("m.user_type='user'").
			Where("m.user_id=?", userId)

		if phone := c.Query("phone"); phone != "" {
			m = m.Where("u.phone=?", phone)
		}

		if startTime, isExist := c.GetQuery("begin_time"); isExist && startTime != "" {
			m = m.Where("m.created_at > ?", startTime)
		}
		if endTime, isExist := c.GetQuery("end_time"); isExist && endTime != "" {
			m = m.Where("m.created_at < ?", endTime)
		}

	}else {
		m = DB.Table("maibao as m").
			Select("es.`phone`,m.`money`,m.`amount` as `maibao`,es.`name`,p.`name` as region, m.user_type,e.`name` as `additional`,m.created_at,m.`type`").
			Joins("LEFT JOIN enterprise_users as eu on eu.id=m.user_id").
			Joins("LEFT JOIN enterprise e on e.id=eu.enterprise_id").
			Joins("LEFT JOIN enterprise_staff es on es.id=eu.staff_id").
			Joins("LEFT JOIN places p on p.id=e.province_id").
			Where("m.user_type='enterprise'").
			Where("m.user_id=?", userId)

		if phone := c.Query("phone"); phone != "" {
			m = m.Where("es.phone=?", phone)
		}
		if startTime, isExist := c.GetQuery("begin_time"); isExist && startTime != "" {
			m = m.Where("m.created_at > ?", startTime)
		}
		if endTime, isExist := c.GetQuery("end_time"); isExist && endTime != "" {
			m = m.Where("m.created_at < ?", endTime)
		}
	}
	var count int
	m.Count(&count)
	paginate := NewPage(page, limit, count)
	m.Limit(paginate.Limit).Offset(paginate.Offset).Scan(&res)
	return &res, &paginate

}
func MaibaoRecharge(userId int, amount float64, userType string) bool {
	tx := DB.Begin()
	var err error
	if userType == "user"{
		tx.Raw("update `users` set `maibao`=`maibao`+?,`maibao_real`=`maibao_real`+? where id=?", amount, amount, userId).Row()
		err = tx.Table("maibao").Create(&struct {
			UserType string
			UserId int
			Amount float64
			Money float64
			Type string
			CreatedAt time.Time
		}{
			UserType:"user",
			UserId:userId,
			Amount:amount,
			Money:0,
			Type:"admin",
			CreatedAt:time.Now(),
		}).Error

	}else{
		err = tx.Raw("update `enterprise` set `maibao`=`maibao`+? where id=?", amount, amount, userId).Error
		err = tx.Table("maibao").Create(&struct {
			UserType string
			UserId int
			Amount float64
			Money float64
			Type string
			CreatedAt time.Time
		}{
			UserType:"enterprise",
			UserId:userId,
			Amount:amount,
			Money:0,
			Type:"admin",
			CreatedAt:time.Now(),
		}).Error
	}
	if err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func MaibaoSpendingLog(page int, limit int, c *gin.Context, userType string) (interface{}, *Paginator){
	var res []*struct{
		Name string `json:"name"`
		Phone string `json:"phone"`
		Additional string `json:"additional"`
		Region string `json:"region"`
		Maibao float64 `json:"maibao"`
		Reg float64 `json:"reg"`
		Gain float64 `json:"gain"`
		Surplus float64 `json:"surplus"`
		Status int `json:"status"`
		UserId int `json:"user_id"`
	}
	var m *gorm.DB
	if userType == "user" {
		m = Table("maibao_bill as mb").Select("mb.`user_id`,c.`user_name` as `name`, u.`phone`, p.`name` as region, u.`maibao_all` as `maibao`, u.`maibao` as `surplus`, r.`amount` as reg, r1.`amount` as `gain`,u.`is_enabled` as `status`,'个人' as additional").
			Joins("left join users as u on u.id=mb.user_id").
			Joins("left join cards c on c.user_id=mb.user_id").
			Joins("left join places p on p.id=c.province").
			Joins("left join (select sum(amount) as amount, user_id from maibao where user_type='user' group by user_id) as r on r.user_id=mb.user_id").
			Joins("left join (select sum(amount) as amount, user_id from maibao_bill where type=1 group by user_id) as r1 on r1.user_id=mb.user_id").
			Where("mb.type=-1").
			Group("mb.user_id")

		if region := c.Query("region"); region != "" {
			m = m.Where("p.name=?", region)
		}
		if status := c.Query("status"); status != "" {
			m = m.Where("u.is_enabled=?", status)
		}
		if phone := c.Query("phone"); phone != "" {
			m = m.Where("u.phone=?", phone)
		}
		if name := c.Query("name"); name != "" {
			m = m.Where("c.user_name=?", name)
		}
	}else{
		m = Table("enterprise_maibao_bill as mb").Select("mb.`enterprise_id` as `user_id`,u.`name`, u.`phone`, p.`name` as region, e.`maibao_all` as `maibao`, e.`maibao` as `surplus`, r.`amount` as reg, r1.`amount` as `gain`,e.`status`,e.`name` as additional").
			Joins("left join enterprise_staff as u on u.id=mb.staff_id").
			Joins("left join enterprise e on e.id=mb.enterprise_id").
			Joins("left join enterprise_users eu on eu.staff_id=mb.staff_id").
			Joins("left join places p on p.id=e.province_id").
			Joins("left join (select sum(amount) as amount, user_id from maibao where user_type='enterprise' group by user_id) as r on r.user_id=mb.staff_id").
			Joins("left join (select sum(amount) as amount, staff_id from enterprise_maibao_bill where type=1 group by enterprise_id) as r1 on r1.staff_id=mb.staff_id").
			Where("mb.type=-1").
			Group("mb.enterprise_id")

		if region := c.Query("region"); region != "" {
			m = m.Where("p.name=?", region)
		}
		if status := c.Query("status"); status != "" {
			m = m.Where("eu.status=?", status)
		}
		if phone := c.Query("phone"); phone != "" {
			m = m.Where("u.phone=?", phone)
		}
		if name := c.Query("name"); name != "" {
			m = m.Where("u.name=?", name)
		}
	}
	var count int
	m.Count(&count)
	paginate := NewPage(page, limit, count)
	m.Limit(paginate.Limit).Offset(paginate.Offset).Scan(&res)
	return &res, &paginate
}

func MaibaoGain(page int, limit int, c *gin.Context, userType string, userId int ,cate int) (interface{}, *Paginator) {
	var res []*struct{
		Name string `json:"name"`
		Phone string `json:"phone"`
		Additional string `json:"additional"`
		Cate int `json:"cate"`
		Title string `json:"title"`
		Amount float64 `json:"amount"`
		CreatedAt time.Time `json:"created_at"`
	}
	var m *gorm.DB
	if userType == "user" {
		m = Table("maibao_bill as mb").Select("c.`user_name` as `name`, u.`phone`, '个人' as `additional`, mb.cate, mb.amount, mb.created_at,mb.title").
			Joins("left join users u on u.id=mb.user_id").
			Joins("left join cards c on c.user_id=mb.user_id").
			Where("mb.user_id=?", userId).
			Where("mb.type=?", cate)

		if phone := c.Query("phone"); phone != "" {
			m = m.Where("u.phone=?", phone)
		}
		if startTime, isExist := c.GetQuery("begin_time"); isExist && startTime != "" {
			m = m.Where("mb.created_at > ?", startTime)
		}
		if endTime, isExist := c.GetQuery("end_time"); isExist && endTime != "" {
			m = m.Where("mb.created_at < ?", endTime)
		}
	}else{
		m = Table("enterprise_maibao_bill as mb").Select("u.`name`, u.`phone`, e.`name` as `additional`, mb.cate, mb.amount, mb.created_at,mb.title").
			Joins("left join enterprise_staff u on u.id=mb.staff_id").
			Joins("left join enterprise e on e.id=mb.enterprise_id").
			Where("mb.staff_id=?", userId).
			Where("mb.type=?", cate)

		if phone := c.Query("phone"); phone != "" {
			m = m.Where("u.phone=?", phone)
		}
		if startTime, isExist := c.GetQuery("begin_time"); isExist && startTime != "" {
			m = m.Where("mb.created_at > ?", startTime)
		}
		if endTime, isExist := c.GetQuery("end_time"); isExist && endTime != "" {
			m = m.Where("mb.created_at < ?", endTime)
		}
	}
	var count int
	m.Count(&count)
	paginate := NewPage(page, limit, count)
	m.Limit(paginate.Limit).Offset(paginate.Offset).Scan(&res)
	return &res, &paginate
}