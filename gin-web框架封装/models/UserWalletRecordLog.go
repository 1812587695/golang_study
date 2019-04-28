package models

import "time"

type UserWalletRecordLog struct {
	ID int `gorm:"primary_key"`
	UserId int
	Type int
	IncomeType int
	RelationId int
	Describe string
	Money float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserWalletRecordLog) TableName() string {
	return "user_wallet_record_log"
}