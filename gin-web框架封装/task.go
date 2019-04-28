package main

import (
	"hytx_manager/models"
	"hytx_manager/pkg/setting"
	"strings"
	"time"
)

func main() {
	//链接数据库
	setting.Setup()
	models.Setup()
	//开始任务
	//所有城市管理
	var m []*struct{
		ID int
		AdminName string
		AdminMobilePhone string
		BankCardNo string
		BankCardName string
		BankName string
		Money float64
	}
	models.DB.Table("place_partners as p").
		Select("p.id,p.admin_name,p.admin_mobile_phone,p.bank_card_no,p.bank_card_name,p.bank_name,IFNULL(bonus_sub.money,0)+IFNULL(o1.agent,0)/2 as money").
		Joins("left join (select user_id, sum(money) as money from user_place_partner_bonus_count where income_type=1 and PERIOD_DIFF(date_format(now(), '%Y%m'),date_format(created_at,'%Y%m')) = 1 group by user_id) as bonus_sub on bonus_sub.user_id=p.user_id").
		Joins("left join (select place_partner_id, sum(agency_fee) as agent from operators where examine=1 and PERIOD_DIFF(date_format(now(), '%Y%m'),date_format(created_at,'%Y%m')) = 1 group by place_partner_id) as o1 on o1.place_partner_id=p.id").
		Scan(&m)

	tx := models.DB.Begin()
	sqlStr := "INSERT INTO `bonus_cash` (`type`,`target_id`,`bank`,`bank_card_no`,`bank_card_name`,`money`,`pay_time`) VALUES "
	var vals []interface{}
	const rowSql = "(?,?,?,?,?,?,?)"
	var inserts []string
	for _, elem := range m {
		if elem.Money > 0 {
			inserts = append(inserts, rowSql)
			vals = append(vals, 1, elem.ID, elem.BankName, elem.BankCardNo, elem.BankCardName, elem.Money, time.Now().Add(25*24*time.Hour))
		}
	}
	sqlStr = sqlStr + strings.Join(inserts, ",")
	var err error
	if vals != nil {
		err = tx.Exec(sqlStr, vals...).Error
	}
	if err != nil {
		tx.Rollback()
		print(err)
	}else {
		tx.Commit()
	}

	var o []*struct{
		ID int
		BankCardNo string
		BankCardName string
		BankName string
		Money float64
	}

	models.DB.Table("operators as p").
		Select("p.id,p.name,p.phone,p.bank_card_no,p.bank_card_name,p.bank_name,bonus_sub.money").
		Joins("left join (select user_id, sum(money) as money from user_operators_bonus_count where income_type=1 and PERIOD_DIFF(date_format(now(), '%Y%m'),date_format(created_at,'%Y%m')) = 1 group by user_id) as bonus_sub on bonus_sub.user_id=p.user_id").
		Scan(&o)

	otx := models.DB.Begin()
	osqlStr := "INSERT INTO `bonus_cash` (`type`,`target_id`,`bank`,`bank_card_no`,`bank_card_name`,`money`,`pay_time`) VALUES "
	var ovals []interface{}
	const orowSql = "(?,?,?,?,?,?,?)"
	var oinserts []string
	for _, elem := range o {
		if elem.Money > 0 {
			oinserts = append(oinserts, orowSql)
			ovals = append(ovals, 2, elem.ID,  elem.BankName, elem.BankCardNo, elem.BankCardName, elem.Money, time.Now().Add(25*24*time.Hour))
		}
	}
	osqlStr = osqlStr + strings.Join(oinserts, ",")
	var oerr error
	if ovals != nil {
		oerr = otx.Exec(osqlStr, ovals...).Error
	}
	if oerr != nil {
		otx.Rollback()
		print(oerr)
	}else {
		otx.Commit()
	}

	var s []*struct{
		ID int
		BankCardNo string
		BankCardName string
		BankName string
		Money float64
	}

	models.DB.Table("salesmen as p").
		Select("p.id,p.name,p.phone,p.bank_card_no,p.bank_card_name,p.bank_name,bonus_sub.money").
		Joins("left join (select user_id, sum(money) as money from user_salesman_bonus_count where income_type=1 and PERIOD_DIFF(date_format(now(), '%Y%m'),date_format(created_at,'%Y%m')) = 1 group by user_id) as bonus_sub on bonus_sub.user_id=p.user_id").
		Scan(&s)

	stx := models.DB.Begin()
	ssqlStr := "INSERT INTO `bonus_cash` (`type`,`target_id`,`bank`,`bank_card_no`,`bank_card_name`,`money`,`pay_time`) VALUES "
	var svals []interface{}
	const srowSql = "(?,?,?,?,?,?,?)"
	var sinserts []string
	for _, elem := range s {
		if elem.Money > 0 {
			sinserts = append(sinserts, srowSql)
			svals = append(svals, 3, elem.ID, elem.BankName, elem.BankCardNo, elem.BankCardName, elem.Money, time.Now().Add(25*24*time.Hour))
		}
	}
	ssqlStr = ssqlStr + strings.Join(sinserts, ",")
	var serr error
	if svals != nil {
		serr = stx.Exec(ssqlStr, svals...).Error
	}
	if serr != nil {
		stx.Rollback()
		print(serr)
	}else {
		stx.Commit()
	}
}
