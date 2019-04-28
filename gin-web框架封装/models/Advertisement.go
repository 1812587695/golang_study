package models

import (
	"time"
)

type Advertisement struct {
	ID int `json:"id"`
	Operator int `json:"operator"`
	Type int `json:"type"`
	Name string `json:"name"`
	Images string `json:"images"`
	Describe string `json:"describe"`
	Url string `json:"url"`
	Cate int `json:"cate"`
	Label int `json:"label"`
	SurvivalBegin time.Time `json:"survival_begin"`
	SurvivalEnd time.Time `json:"survival_end"`
	Status int `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Advertisement) TableName() string {
	return "advertisement"
}

func AdList(page int, pageSize int, status string, beginTime string, endTime string) (interface{}, Paginator){
	var res []*struct{
		Advertisement
		RestTime int64 `json:"rest_time"`
		CateName string `json:"cate_name"`
		LabelName string `json:"label_name"`
	}
	model:= DB.Table("advertisement as a").
		Select("a.*,ac.title as cate_name,al.name as label_name").
		Joins("left join article_categories as ac on ac.id=a.cate").
		Joins("left join article_label as al on al.id=a.label")
	if status != "" {
		model = model.Where("status=?", status)
	}
	if beginTime != "" {
		model = model.Where("created_at > ?", beginTime)
	}
	if endTime != "" {
		model = model.Where("created_at < ?", endTime)
	}
	var count int
	model.Count(&count)
	paginate := NewPage(page, pageSize, count)
	err := model.Offset(paginate.Offset).Limit(paginate.Limit).Find(&res).Error
	if err != nil {
		panic(err)
	}
	for _, i := range res{
		//结束时间在当前之前
		if i.SurvivalEnd.Before(time.Now()) {
			i.RestTime = 0
		}else{
			i.RestTime = i.SurvivalEnd.Unix() - time.Now().Unix()
		}
	}
	return res, paginate
}

func (a *Advertisement) Add() error {
	return DB.Model(&Advertisement{}).Create(a).Error
}

func GetAdById(id string) *Advertisement {
	var res Advertisement

	DB.Model(&Advertisement{}).Where("id=?", id).First(&res)
	return &res
}
