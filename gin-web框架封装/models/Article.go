package models

import (
	"time"
)

type Article struct {
	ID                int       `gorm:"primary_key" json:"id"`
	Author            int       `json:"author"`
	Title             string    `json:"title"`
	SubTitle          string    `json:"sub_title"`
	LabelId           string    `json:"label_id"`
	Content           string    `json:"content"`
	Cover             string    `json:"cover"`
	AdvertisementId   int       `json:"advertisement_id"`
	IsAd              int       `json:"is_ad"`
	ArticleCategoryId int       `json:"article_category_id"`
	Sort              int       `json:"sort"`
	Creator           int       `json:"creator"`
	Pageviews         int       `json:"pageviews"`
	ActualPageviews   int       `json:"actual_pageviews"`
	CommentNum        int       `json:"comment_num"`
	LastCommentAt     time.Time `json:"last_comment_at"`
	Status            int       `json:"status"`
	ShareNum          int       `json:"share_num"`
	RecommendNum      int       `json:"recommend_num"`
	CollectNum        int       `json:"collect_num"`
	OperateStatus     int       `json:"operate_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsDeleted         int       `json:"is_deleted"`
}

type FindArticlesList struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Cover           string    `json:"cover"`
	CategoryName    string    `json:"category_name"`
	LabelName       string    `json:"label_name"`
	CreatedAt       time.Time `json:"created_at"`
	IsAd            int       `json:"is_ad"`
	AdvertisementId int       `json:"advertisement_id"`
	Status          int       `json:"status"`
}

func ArticleExamineCount() int {
	model := DB.Table("articles").Where("status=?", 1)
	var count int
	model.Count(&count)
	return count
}

// 查找
func GetArticles(status string, article_category_id int, operation string, begin_time string, end_time string, keyword  string,page int, pageSize int) ([]*FindArticlesList, *Paginator) {
	var data []*FindArticlesList
	item := DB.Table("articles as  t1 ").Select("t1.id, t1.title, t1.cover, t2.title as category_name, t3.name as  label_name , t1.created_at, t1.is_ad,t1.advertisement_id, t1.status").Joins("left join article_categories as t2 on t1.article_category_id = t2.id ").Joins("left join article_label as  t3 on t1.label_id = t3.id")
	if status != "" {
		item = item.Where("t1.status=?", status)
	}
	if article_category_id != 0 {
		item = item.Where("t1.article_category_id=?", article_category_id)
	}

	if begin_time != "" {
		item = item.Where("t1.created_at > ?", begin_time)
	}
	if end_time != "" {
		item = item.Where("t1.created_at < ?", end_time)
	}
	if keyword != "" {
		item = item.Where("t1.title like ?", "%"+keyword+"%")
	}
	item = item.Where("t1.is_deleted=0")
	item = item.Where("isNull(t1.deleted_at)")
	if operation != "" {
		item = item.Order(operation + " desc")
	} else {
		item = item.Order("id" + " desc")
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

// 更换广告
func UpdateArticleAdvertisementId(id int, advertisement_id int) error {
	var data = Article{AdvertisementId: advertisement_id, IsAd: 1}
	err := DB.Table("articles  ").Where("id=?", id).Updates(data).Error

	if err != nil {
		return err
	}
	return nil
}

// 查找单个
func GetArticleById(id int) (*Article, error) {
	var data Article
	err := DB.Table("articles").Where("id=?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// 编辑
func UpdateArticle(data map[string]interface{}) error {
	err := DB.Table("articles").Where("id=?", data["id"]).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除
func DeletedArticle(id int) error {
	err := DB.Delete(Article{}, "id=?", id).Error
	if err != nil {
		return err
	}
	return nil
}
