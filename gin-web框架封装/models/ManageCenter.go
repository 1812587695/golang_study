package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//管理中心
type ManageCenter struct {
	ID               int       `json:"id"`
	PID              int       `json:"pid" gorm:"column:pid"`
	AID              int       `gorm:"column:aid"`
	UserId           int       `json:"user_id"`
	Province         int       `json:"province"`
	City             int       `json:"city"`
	Area             int       `json:"area"`
	Title            string    `json:"title"`
	AdminName        string    `json:"username"`
	AdminMobilePhone string    `json:"phone"`
	IdCard           string    `json:"id_card"`
	IdCardFront      string    `json:"id_card_front"`
	IdCardBack       string    `json:"id_card_back"`
	BankCardNo       string    `json:"bank_card_no"`
	BankCardName     string    `json:"bank_card_name"`
	BankName         string    `json:"bank_name"`
	BankBranch       string    `json:"bank_branch"`
	OperatorCount    int       `json:"operator_count"`
	SalesmanCount    int       `json:"salesman_count"`
	UserCount        int       `json:"user_count"`
	VipCount         int       `json:"vip_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Status           int       `json:"status"`
	OperatorsNum     int       `json:"operators_num"`
	BusinessLicense  string    `json:"business_license"`
}

func (m *ManageCenter) TableName() string {
	return "place_partners"
}

func (m *ManageCenter) Insert() bool {
	err := DB.Create(m).Error
	if err != nil {
		return false
	}
	return true
}

func (m *ManageCenter) Update(id string) bool {
	err := DB.Model(m).Where("id=?", id).Update(m).Error
	if err != nil {
		return false
	}
	return true
}
func ManageCenterList(page int, limit int, c *gin.Context) (interface{}, *Paginator) {
	var res []*struct {
		ID           int       `json:"id"`
		UserId int `json:"-"`
		Title        string    `json:"title"`
		Name         string    `json:"name"`
		Phone        string    `json:"phone"`
		VipNum       int       `json:"vipNum"`
		OperatorsNum int       `json:"operatorsNum"`
		SalesmanNum  int       `json:"salesmanNum"`
		UserNum      int       `json:"userNum"`
		CreatedAt    time.Time `json:"createdAt"`
		State        int       `json:"state"`
	}
	m := DB.Table("place_partners as p").
		Select("p.id,p.title,p.user_id, p.admin_name as name, p.admin_mobile_phone as phone,o.operators_num,s.salesman_num,u.user_num,p.created_at,p.status as state").
		Joins("left join (select count(id) as operators_num,place_partner_id from users where type=2 group by place_partner_id) as o on o.place_partner_id=p.user_id").
		Joins("left join (select count(id) as salesman_num,place_partner_id from users where type=3 group by place_partner_id) as s on s.place_partner_id=p.user_id").
		Joins("left join (select count(id) as user_num,place_partner_id from users where type=4 group by place_partner_id) as u on u.place_partner_id=p.user_id")
	if state, isExist := c.GetQuery("state"); isExist && state != "" {
		m = m.Where("status=?", state)
	}
	if area, isExist := c.GetQuery("area"); isExist && area != "" {
		var id int
		err := DB.Table("places").Select("`id`").Where("p.name like ?", "%"+area+"%").Where("p.level=1").Row().Scan(&id)
		if err != nil {
			fmt.Println(err)
		}
		m = m.Where("p.province=?", id)
	}
	if startTime, isExist := c.GetQuery("begin_time"); isExist && startTime != "" {
		m = m.Where("p.created_at > ?", startTime)
	}
	if endTime, isExist := c.GetQuery("end_time"); isExist && endTime != "" {
		m = m.Where("p.created_at < ?", endTime)
	}
	if phone, isExist := c.GetQuery("phone"); isExist && phone != "" {
		m = m.Where("p.admin_mobile_phone=?", phone)
	}
	var count int
	m.Count(&count)
	paginate := NewPage(page, limit, count)
	m.Scan(&res)
	for _, i := range res {
		var n int
		DB.Table("user_grade").Where("user_id in (SELECT id from users where place_partner_id=?)", i.UserId).Count(&n)
		i.VipNum = n
	}
	return res, &paginate
}
func GetManageCenterById(id int) *ManageCenter {
	var mo ManageCenter
	DB.Model(ManageCenter{}).Where("id=?", id).First(&mo)
	return &mo
}

func (m *ManageCenter) Recharges(c *gin.Context, page int, limit int) (interface{}, *Paginator){
	var res []*struct{
		ID int `json:"id"`
		Title string `json:"title"`
		UserId int `json:"user_id"`
		AllMoney float64 `json:"all_money"`
		LastMonthMoney float64 `json:"last_month_money"`
		CurrentMoney float64 `json:"current_money"`
		Admin string `json:"admin"`
		Phone string `json:"phone"`
	}
	model := DB.Table("place_partners as p").
		Select("p.id,p.title,p.user_id,p.admin_name as admin,p.admin_mobile_phone as phone,IFNULL(bonus.money,0)+IFNULL(o.agent,0)/2 as all_money,IFNULL(bonus_sub.money,0)+IFNULL(o1.agent,0)/2 as last_month_money,IFNULL(bonus_to.money,0)+IFNULL(o2.agent,0)/2 as current_money").
		Joins("left join (select user_id, sum(money) as money from user_place_partner_bonus_count where income_type=1 group by user_id) as bonus on bonus.user_id=p.user_id").
		Joins("left join (select user_id, sum(money) as money from user_place_partner_bonus_count where income_type=1 and PERIOD_DIFF(date_format(now(), '%Y%m'),date_format(created_at,'%Y%m')) = 1 group by user_id) as bonus_sub on bonus_sub.user_id=p.user_id").
		Joins("left join (select user_id, sum(money) as money from user_place_partner_bonus_count where income_type=1 and DATE_FORMAT(created_at,'%Y%m')=DATE_FORMAT(CURDATE(), '%Y%m') group by user_id) as bonus_to on bonus_to.user_id=p.user_id").
		Joins("left join (select place_partner_id, sum(agency_fee) as agent from operators where examine=1 group by place_partner_id) as o on o.place_partner_id=p.id").
		Joins("left join (select place_partner_id, sum(agency_fee) as agent from operators where examine=1 and PERIOD_DIFF(date_format(now(), '%Y%m'),date_format(created_at,'%Y%m')) = 1 group by place_partner_id) as o1 on o1.place_partner_id=p.id").
		Joins("left join (select place_partner_id, sum(agency_fee) as agent from operators where examine=1 and DATE_FORMAT(created_at,'%Y%m')=DATE_FORMAT(CURDATE(), '%Y%m') group by place_partner_id) as o2 on o2.place_partner_id=p.id")

	var count int
	model.Count(&count)
	paginate := NewPage(page, limit, count)
	model.Scan(&res)

	return res, &paginate
}

func (m *ManageCenter) RechargesDetail(c *gin.Context, page int, limit int) interface{}  {
	id := c.Query("id")
	var m1 ManageCenter
	DB.Model(m).Where("id=?", id).First(&m1)
	var total float64
	var vipTotal float64
	var agentFee float64
	var eTotal float64
	var vipToday float64
	var agentFeeToday float64
	var lastMonthTotal float64
	var lastMonthAgent float64
	var currMonthTotal float64
	var currMonthAgent float64
	_ = Table("user_place_partner_bonus_count").Select("sum(money)").Where("user_id=?", m1.UserId).Where("income_type=1").Row().Scan(&total)
	_ = Table("user_place_partner_bonus_count").Select("sum(money)").Where("income_type=1").Where("user_id=?", m1.UserId).Where("type=1").Row().Scan(&vipTotal)
	_ = Table("user_place_partner_bonus_count").Select("sum(money)").Where("income_type=1").Where("user_id=?", m1.UserId).Where("type=2").Row().Scan(&eTotal)
	_ = Table("operators").Select("sum(agency_fee)").Where("examine=1").Where("place_partner_id=?", id).Row().Scan(&agentFee)
	_ = Table("user_place_partner_bonus_count").Select("sum(money)").Where("user_id=?", m1.UserId).Where("income_type=1").Where("type=1").Where("to_days(created_at)=to_days(now())").Row().Scan(&vipToday)
	_ = Table("operators").Select("sum(agency_fee)").Where("place_partner_id=?", id).Where("examine=1").Where("to_days(examined_at)=to_days(now())").Row().Scan(&agentFeeToday)
	//上月奖励金
	_ = Table("user_place_partner_bonus_count").Select("sum(money)").Where("income_type=1").Where("user_id=?", m1.UserId).Where("PERIOD_DIFF(date_format(now(), '%Y%m'),date_format(created_at,'%Y%m')) = 1").Row().Scan(&lastMonthTotal)
	_ = Table("operators").Select("sum(agency_fee)").Where("place_partner_id=?", id).Where("examine=1").Where("PERIOD_DIFF(date_format(now(), '%Y%m'),date_format(created_at,'%Y%m')) = 1").Row().Scan(&lastMonthAgent)
	lastMonthTotal = lastMonthTotal + (lastMonthAgent / 2)
	//当月
	_ = Table("user_place_partner_bonus_count").Select("sum(money)").Where("income_type=1").Where("user_id=?", m1.UserId).Where("DATE_FORMAT(created_at,'%Y%m')=DATE_FORMAT(CURDATE(), '%Y%m')").Row().Scan(&currMonthTotal)
	_ = Table("operators").Select("sum(agency_fee)").Where("place_partner_id=?", id).Where("examine=1").Where("DATE_FORMAT(created_at,'%Y%m')=DATE_FORMAT(CURDATE(), '%Y%m')").Row().Scan(&currMonthAgent)
	currMonthTotal = currMonthTotal + (currMonthAgent / 2)
	total = total + vipTotal + agentFee
	return map[string]float64{
		"total":total,
		"vip_total": vipTotal,
		"agent_total": agentFee / 2,
		"e_total" : eTotal,
		"vip_today": vipToday,
		"agent_today": agentFeeToday / 2,
		"last_month_total": lastMonthTotal,
		"curr_month_total": currMonthTotal,
	}
}
func (m *ManageCenter) RechargesVipList(c *gin.Context, page int, limit int) (interface{}, *Paginator)  {
	id := c.Query("id")
	var m1 ManageCenter
	DB.Model(m).Where("id=?", id).First(&m1)
	var res []*struct{
		UserName string	`json:"user_name"`
		Phone string `json:"phone"`
		Relation string `json:"relation"`
		PayTime time.Time `json:"pay_time"`
		VipFeeConfigComment string `json:"vip_level"`
		Money float64 `json:"money"`
		TMoney float64 `json:"t_money"`
		TTime time.Time `json:"t_time"`
	}
	model := Table("user_commission_details as ucd").
		Select("c.user_name,u.phone,ucd.relation,ur.pay_time,ur.vip_fee_config_comment,ur.money,ucd.money as t_money,ucd.created_at as t_time").
		Joins("left join user_recharges as ur on ur.out_trade_no=ucd.out_trade_no").
		Where("ucd.relation_id=?", m1.UserId).
		Joins("left join users u on u.id=ucd.user_id").
		Joins("left join cards c on c.user_id=ucd.user_id")

	if startTime, isExist := c.GetQuery("begin_time"); isExist && startTime != "" {
		model = model.Where("ucd.created_at > ?", startTime)
	}
	if endTime, isExist := c.GetQuery("end_time"); isExist && endTime != "" {
		model = model.Where("ucd.created_at < ?", endTime)
	}
	if phone, isExist := c.GetQuery("phone"); isExist && phone != "" {
		model = model.Where("u.phone=?", phone)
	}
	if vip, isExist := c.GetQuery("vip"); isExist && vip != "" {
		model = model.Where("ur.vip_fee_config_comment=?", vip)
	}

	var count int
	model.Count(&count)
	paginate := NewPage(page, limit, count)
	model.Offset(paginate.Offset).Limit(paginate.Limit).Scan(&res)
	return &res, &paginate
}

func (m *ManageCenter) RechargesAgentList(c *gin.Context, page int, limit int) (interface{}, *Paginator) {
	id := c.Query("id")
	var res []*struct{
		Title string `json:"title"`
		Name string `json:"name"`
		Phone string `json:"phone"`
		AgentMoney float64 `json:"agent_money"`
		ExaminedAt time.Time `json:"examined_at"`
		Money float64 `json:"money"`
	}
	model := Table("operators").Select("title,name,phone,agency_fee as agent_money,examined_at,(agency_fee/2) as money").Where("place_partner_id=?", id).Where("examine=1")
	if startTime, isExist := c.GetQuery("begin_time"); isExist && startTime != "" {
		model = model.Where("created_at > ?", startTime)
	}
	if endTime, isExist := c.GetQuery("end_time"); isExist && endTime != "" {
		model = model.Where("created_at < ?", endTime)
	}
	if phone, isExist := c.GetQuery("phone"); isExist && phone != "" {
		model = model.Where("phone=?", phone)
	}

	var count int
	model.Count(&count)
	paginate := NewPage(page, limit, count)
	model.Offset(paginate.Offset).Limit(paginate.Limit).Scan(&res)
	return &res, &paginate
}

func (m *ManageCenter) RechargesEList(c *gin.Context, page int, limit int) (interface{}, *Paginator) {
	id := c.Query("id")
	var m1 ManageCenter
	DB.Model(m).Where("id=?", id).First(&m1)
	var res []*struct{
		Name string `json:"name"`
		User string `json:"user"`
		Phone string `json:"phone"`
		Money float64 `json:"money"`
		RMoney float64 `json:"r_money"`
		PayTime time.Time `json:"pay_time"`
	}
	model := Table("user_place_partner_bonus_count as upbc").
		Select("e.name,es.name as user,es.phone,er.money,upbc.money as r_money, er.pay_time").
		Joins("left join enterprise_recharge_fenxiao_queue as eq on eq.id=upbc.relation_id").
		Joins("left join enterprise_recharges er on er.out_trade_no=eq.out_trade_no").
		Joins("left join enterprise e on e.id=er.enterprise_id").
		Joins("left join enterprise_staff es on es.id=er.staff_id").
		Where("upbc.user_id=?",m1.UserId)

	if startTime, isExist := c.GetQuery("begin_time"); isExist && startTime != "" {
		model = model.Where("upbc.created_at > ?", startTime)
	}
	if endTime, isExist := c.GetQuery("end_time"); isExist && endTime != "" {
		model = model.Where("upbc.created_at < ?", endTime)
	}
	if phone, isExist := c.GetQuery("phone"); isExist && phone != "" {
		model = model.Where("es.phone=?", phone)
	}

	var count int
	model.Count(&count)
	paginate := NewPage(page, limit, count)
	model.Offset(paginate.Offset).Limit(paginate.Limit).Scan(&res)
	return &res, &paginate
}

func (m *ManageCenter) Cash(c *gin.Context, page int, limit int) (interface{}, *Paginator) {
	id := c.Query("id")
	var res []*struct{
		ID int `json:"id"`
		Bank string `json:"bank"`
		BankCardName string `json:"bank_card_name"`
		BankCardNo string `json:"bank_card_no"`
		Money float64 `json:"money"`
		PayTime time.Time `json:"pay_time"`
		Status int `json:"status"`
	}
	model := Table("bonus_cash").Where("target_id=?", id).Where("type=1")

	var count int
	model.Count(&count)
	paginate := NewPage(page, limit, count)
	model.Offset(paginate.Offset).Limit(paginate.Limit).Scan(&res)
	return &res, &paginate
}