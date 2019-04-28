package models

import "time"

type Cards struct {
	ID                       int       `gorm:"primary_key" json:"id"`
	UserId                   int       `json:"user_id"`
	HeadPicture              string    `json:"head_picture"`
	UserName                 string    `json:"user_name"`
	Motto                    string    `json:"motto"`
	MobilePhone              string    `json:"mobile_phone"`
	Telephone                string    `json:"telephone"`
	Company                  string    `json:"company"`
	Industry                 string    `json:"industry"`
	Profession               string    `json:"profession"`
	Sex                      int       `json:"sex"`
	Birthday                 string    `json:"birthday"`
	WechatCode               string    `json:"wechat_code"`
	Qq                       string    `json:"qq"`
	Email                    string    `json:"email"`
	Website                  string    `json:"website"`
	Province                 string    `json:"province"`
	City                     string    `json:"city"`
	Area                     string    `json:"area"`
	AddressDetail            string    `json:"address_detail"`
	Popularity               int       `json:"popularity"`
	Saved                    int       `json:"saved"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	PersonalPageTemplateUser []*PersonalPageTemplateUserList
	MemberMarketing          []*MemberMarketingList
	VipName                  string    `json:"vip_name"`
}

type PersonalPageTemplateUserList struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MemberMarketingList struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MemberInfoResult struct {
	Id         int    `json:"id"`
	UserName   string `json:"user_name"`
	Phone      string `json:"phone"`
	Company    string `json:"company"`
	Industry   string `json:"industry"`
	Profession string `json:"profession"`
	Province   string `json:"province"`
	IsEnabled  int    `json:"is_enabled"`
	VipName    string `json:"vip_name"`
	CreatedAt                time.Time `json:"created_at"`
}

type MemberInfoList struct {
	Id int `json:"id"`
	Phone string `json:"phone"`
	UserName   string `json:"user_name"`
	Company    string `json:"company"`
}

type PersonalPageTemplateUser struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	TemplateId int    `json:"template_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsDefault  int    `json:"is_default"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Popularity int    `json:"popularity"`
	Cate       int    `json:"cate"`
	Components string `json:"components"`
	Type       int    `json:"type"`
	Background string `json:"background"`
}

type UserMarketing struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Contents  string `json:"contents"`
	IsDeleted int    `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetMembers(is_enabled int, begin_time string, end_time string, company_name string, is_grade int,  page int, pageSize int) ([]*MemberInfoResult, *Paginator) {
	var data []*MemberInfoResult
	item := DB.Table("users as t1").
		Select("t1.id, t2.user_name, t1.phone, t2.company, t2.industry, t2.profession, t2.province, t1.is_enabled, t4.name as vip_name, t1.created_at").
		Joins("left join cards as t2 on t1.id=t2.user_id ").
		Joins("left join user_grade as t3 on t1.id= t3.user_id").
		Joins("left join vip_fee_config as t4 on t3.vip_fee_config_id=t4.id").
		Where("t1.type=4").
		Where("t1.phone <> ''").
		Order("t1.created_at desc",true).
		Find(&data)

	if is_enabled != 0 {
		if is_enabled == 1 {
			item = item.Where("t1.is_enabled=?", 0)
		} else {
			item = item.Where("t1.is_enabled=?", 1)
		}
	}
	if begin_time != "" {
		item = item.Where("t1.created_at > ?", begin_time)
	}
	if end_time != "" {
		item = item.Where("t1.created_at < ?", end_time)
	}

	if company_name != "" {
		item = item.Where("t1.phone like ?", "%"+company_name+"%")
	}

	if is_grade != 0 {
		item = item.Where("t4.id=?", is_grade)
	}
	var count int
	item.Count(&count)
	paginate := NewPage(page, pageSize, count)
	err := item.Offset(paginate.Offset).Limit(paginate.Limit).Find(&data).Error
	if err != nil {
		panic(err)
	}
	return data, &paginate

}

func GetMemberInfoById(user_id int) (map[int]interface{}) {
	var cards []*Cards
	data := make(map[int]interface{})
	DB.Table("cards as b1 ").Select("b1.* , b3.name as vip_name").Joins("left join user_grade as b2 on b1.user_id=b2.user_id").Joins("left join vip_fee_config as b3 on b2.vip_fee_config_id= b3.id").Joins("").Where("b1.user_id=?", user_id).Find(&cards)
	for i := 0; i < len(cards); i++ {
		var personalPageTemplate []*PersonalPageTemplateUserList
		DB.Table("personal_page_template_user as t1").Select("t1.id, t2.name").Joins("left join personal_page_template as t2 on t1.template_id = t2.id").Where("t1.user_id=?", user_id).Find(&personalPageTemplate)
		cards[i].PersonalPageTemplateUser = personalPageTemplate
		var marketing []*MemberMarketingList
		DB.Table("user_marketing").Select("id, title as name").Where("user_id=?", user_id).Find(&marketing)
		cards[i].MemberMarketing = marketing
		data[i] = cards[i]

	}

	return data
}

func GetPersonalPageTemplateUser(id int) (*PersonalPageTemplateUser, error) {
	var data PersonalPageTemplateUser
	err := DB.Table("personal_page_template_user").Where("user_id=?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetUserMarketing(id int) (*UserMarketing, error) {
	var data UserMarketing
	err := DB.Table("user_marketing").Where("user_id=?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
