package models

import (
	"errors"
	"hytx_manager/pkg/util"
	"strings"
	"time"

	"github.com/Unknwon/com"
)

type VipCode struct {
	Id           int       `json:"id"`
	VipCode      string    `json:"vip_code"`
	Days         int       `json:"days"`
	IsUsed       int       `json:"is_used"`
	CreateUserId int       `json:"create_user_id"`
	EnterpriseId int       `json:"enterprise_id"`
	StaffId      int       `json:"staff_id"`
	CreateAt     time.Time `json:"create_at"`
	UsedAt       time.Time `json:"used_at"`
}

func AddVipCode(data map[string]interface{}) error {
	vipCode := makeVipCode()
	m := DB.Table("enterprise_vip_code").Create(&VipCode{
		VipCode:      vipCode,
		Days:         data["days"].(int),
		CreateUserId: data["adminId"].(int),
		CreateAt:     time.Now(),
	})
	if m.Error != nil {
		return errors.New(m.Error.Error())
	}

	return nil
}
func GetVipCodeList(page, limit int, code string) (interface{}, *Paginator) {
	var res []*VipCode
	m := DB.Table("enterprise_vip_code").Order("id DESC")

	if code != "" {
		m = m.Where("vip_code = ?", code)
	}
	var count int
	m.Count(&count)
	paginate := NewPage(page, limit, count)

	err := m.Offset(paginate.Offset).Limit(paginate.Limit).Find(&res).Error
	if err != nil {
		return nil, nil
	}

	return res, &paginate
}

func makeVipCode() string {
	string := strings.ToUpper(util.RandMD5(com.ToStr(time.Now().UnixNano())))

	return string[8:24]
}
