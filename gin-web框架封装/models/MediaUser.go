package models

import (
	"time"
)

type MediaUser struct {
	ID               int       `gorm:"primary_key" json:"id"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Identity         int       `json:"identity"`
	Status           int       `json:"status"`
	AuthenticationAt time.Time `json:"authentication_at"`
	Score            int       `json:"score"`
	Remarks          string    `json:"remarks"`
	BlockAt          time.Time `json:"block_at"`
	UnsealAt         time.Time `json:"unseal_at"`
	IsEnabled        int       `json:"is_enabled"`
	CreatedAt        time.Time `json:"created_at"`
	UpdateAt         time.Time `json:"update_at"`
}

type MediaUserResult struct {
	Id        int       `json:"id"`
	Phone     string    `json:"phone"`
	Account   string    `json:"account"`
	Status    int       `json:"status"`
	Identity  int       `json:"identity"`
	IsEnabled int       `json:"is_enabled"`
	CreatedAt time.Time `json:"created_at"`
}

type MediaUserInfoByIdResult struct {
	Id               int       `json:"id"`
	Account          string    `json:"account"`
	Introduce        string    `json:"introduce"`
	Avatar           string    `json:"avatar"`
	OperatorsName    string    `json:"operators_name"`
	OperatorsIdNo    string    `json:"operators_id_no"`
	OperatorsIdPhoto string    `json:"operators_id_photo"`
	Address          string    `json:"address"`
	Email            string    `json:"email"`
	Status           int       `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

type MediaUserInfoInstitutionByIdResult struct {
	Id                      int       `json:"id"`
	Account                 string    `json:"account"`
	Introduce               string    `json:"introduce"`
	Avatar                  string    `json:"avatar"`
	OperatorsName           string    `json:"operators_name"`
	OperatorsIdNo           string    `json:"operators_id_no"`
	OperatorsIdPhoto        string    `json:"operators_id_photo"`
	Address                 string    `json:"address"`
	Email                   string    `json:"email"`
	Status                  int       `json:"status"`
	OtherAptitude           string    `json:"other_aptitude"`
	AuxData                 string    `json:"aux_data"`
	InstitutionName         string    `json:"institution_name"`
	InstitutionLicense      string    `json:"institution_license"`
	InstitutionConfirmation string    `json:"institution_confirmation"`
	InstitutionWebSite      string    `json:"institution_web_site"`
	InstitutionAddress      string    `json:"institution_address"`
	CreatedAt               time.Time `json:"created_at"`
}

func (a *MediaUser) TableName() string {
	return "media_users"
}

func MediaUserExamineCount() int {
	var count int
	DB.Table("media_users").Where("status=?", 1).Count(&count)
	return count
}

func GetMediaUsersAll(status int, identity int, begin_time string, end_time string, phone string, page int, pageSize int) ([]*MediaUserResult, *Paginator) {
	var data []*MediaUserResult

	item := DB.Table("media_users t1").Select("t1.id, t1.phone, t1.identity, t1.status, t2.account, t1.created_at, t1.is_enabled").Joins("left join media_user_info t2 on t1.id = t2.media_user_id").Find(&data)
	if status != 0 {
		item = item.Where("t1.status=?", status)
	}
	if identity != 0 {
		item = item.Where("t1.identity=?", identity)
	}
	if phone != "" {
		item = item.Where("t1.phone=?", phone)
	}

	if begin_time != "" {
		item = item.Where("t1.created_at > ?", begin_time)
	}
	if end_time != "" {
		item = item.Where("t1.created_at < ?", end_time)
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

//通过id 获取自媒体用户信息(个人)
func GetMediaUserInfoById(id int) ([]*MediaUserInfoByIdResult, error) {
	var data []*MediaUserInfoByIdResult
	err := DB.Table("media_users t1 ").Select("t1.id, t2.account , t2.introduce, t2.avatar,t2.operators_name, t2.operators_id_no, t2.operators_id_photo,t2.address, t2.email, t1.status, t1.created_at").Joins("left join media_user_info t2 on t1.id=t2.media_user_id").Where("t1.id=?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 通过id获取自媒体机构用户信息

func GetMediaUserInfoInstitution(id int) ([]*MediaUserInfoInstitutionByIdResult, error) {
	var data []*MediaUserInfoInstitutionByIdResult
	err := DB.Table("media_users t1").Select("t1.id, t2.account, t2.introduce, t2.avatar,t2.operators_name, t2.operators_id_no, t2.operators_id_photo,t2.address, t2.email, t1.status, t3.other_aptitude, t3.aux_data,t3.institution_name,t3.institution_license,t3.institution_confirmation,t3.institution_web_site, t3.institution_address, t1.created_at").Joins("left join media_user_info t2 on t1.id=t2.media_user_id").Joins("left join media_user_institution t3 on t1.id=t3.media_user_id").Where("t1.id=?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UpdateMediaUserStatus(id int, status int, remarks string) (error) {
	err := DB.Table("media_users").Where("id=?", id).Updates(map[string]interface{}{
		"status":            status,
		"authentication_at": time.Now(),
		"remarks":           remarks,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
