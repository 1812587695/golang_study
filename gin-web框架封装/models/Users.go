package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Users struct {
	ID                 int       `gorm:"primary_key"json:"id"`
	Account            string    `gorm:"default:''"json:"account"`
	Phone              string    `gorm:"default:''"json:"phone"`
	Password           string    `gorm:"default:''"json:"password"`
	Email              string    `gorm:"default:''"json:"email"`
	Type               int       `gorm:"default:4"json:"type"`
	Money              float64   `gorm:"default:0"json:"money"`
	PaymentPassword    string    `gorm:"default:''"json:"payment_password"`
	AvailableAmount    float64   `gorm:"column:availableAmount;default:''"json:"available_amount"`
	IsEnabled          int       `gorm:"default:0"json:"is_enabled"`
	LoggedIp           string    `gorm:"default:''"json:"logged_ip"`
	LoggedAt           time.Time `gorm:"default:0"json:"logged_at"`
	AgencyFee          float64   `gorm:"default:0"json:"agency_fee"`
	CurrentCommission  float64   `gorm:"column:currentCommission;default:0"json:"current_commission"`
	ChangeCommission   float64   `gorm:"column:changeCommission;default:0"json:"change_commission"`
	QrCode             string    `gorm:"default:''"json:"qr_code"`
	CreatedAt          time.Time `gorm:"default:0"json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:0"json:"updated_at"`
	Pid                int       `json:"pid"`
	SalesmanId         int       `json:"salesman_id"`
	OperatorId         int       `json:"operator_id"`
	PlacePartnerId     int       `json:"place_partner_id"`
	IsChangeGrade      int       `json:"is_change_grade"`
	IsFaseDistribution int       `json:"is_fase_distribution"`
	EnterpriseUserId   int       `json:"enterprise_user_id"`
}

func (u *Users) TableName() string {
	return "users"
}

func (u *Users) Insert() bool {
	err := DB.Create(u).Error
	if err != nil {
		return false
	}
	return true
}

func (u *Users) InsertDB(db *gorm.DB) bool {
	err := db.Create(u).Error
	if err != nil {
		return false
	}
	return true
}

func (u *Users) Update() bool {
	err := DB.Model(u).Where("id=?", u.ID).Update(u).Error
	if err != nil {
		return false
	}
	return true
}

// 获取用户一条信息
func (u *Users) GetUserInfoById() (*Users, error) {
	var user Users

	err := DB.Model(u).Where("id=?", u.ID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 获取用户一条信息
func (u *Users) GetUserInfoByPhone() (*Users, error) {
	var user Users

	err := DB.Model(u).Where("phone=?", u.Phone).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (e *Users) Count() int {

	count := 0
	DB.Model(Users{}).Where("enterprise_user_id=?", e.EnterpriseUserId).Count(&count)

	return count
}
