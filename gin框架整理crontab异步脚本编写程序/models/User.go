package models

type UserInterPushInfo struct {
	ID           int    `gorm:"primary_key"json:"id"`
	UserId       int    `json:"user_id"`
	MaterialId   int    `json:"material_id"`
	MaterialType string `json:"material_type"`
	ProvinceId   int    `json:"province_id"`
	Province     string `json:"province"`
	Num          int    `json:"num"`
	UseNum       int    `json:"use_num"`
	IsEnabled    int    `json:"is_enabled"`
}

/**
	用户互推信息表
 */
func (m *UserInterPushInfo) TableName() string {
	return "user_interpush_info"
}

// 获取最后一个id
func (i *UserInterPushInfo) GetLastId() (*UserInterPushInfo, error) {
	var in UserInterPushInfo

	err := DB.Model(i).Select("id").Order("id desc").Limit("1").Find(&in).Error
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (i *UserInterPushInfo) Get() (*UserInterPushInfo, error) {
	var in UserInterPushInfo

	err := DB.Model(i).Where("id = ? and is_enabled = ?", i.ID, i.IsEnabled).Find(&in).Error
	if err != nil {
		return nil, err
	}

	return &in, nil
}
