package models

import "time"

type VipFeeConfig struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	Days        int       `json:"days"`
	Type        int       `json:"type"`
	Months      int       `json:"months"`
	Money       float64   `json:"money"`
	LastMoney   float64   `json:"last_money"`
	AutoCharge  int       `json:"auto_charge"`
	IconPath    string    `json:"icon_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
	Description string    `json:"description"`
}

type VipFeeConfigSelectResult struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func GetVipFeeConfig() ([]*VipFeeConfig, error) {
	var data []*VipFeeConfig
	err := DB.Table("vip_fee_config").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetVipFeeConfigSelectResult()  ([]*VipFeeConfigSelectResult, error){
	var data []*VipFeeConfigSelectResult
	err := DB.Table("vip_fee_config").Select("id, name").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
// 获取会员充值等级表一条信息
func (v *VipFeeConfig) GetVipFeeConfigInfoById() (*VipFeeConfig, error) {
	var vip VipFeeConfig

	err := DB.Model(v).Where("id=?", v.ID).Find(&vip).Error
	if err != nil {
		return nil, err
	}
	return &vip, nil
}
