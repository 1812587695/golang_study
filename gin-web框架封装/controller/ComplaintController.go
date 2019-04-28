package controller

import (
	"github.com/gin-gonic/gin"
	"hytx_manager/middleware/auth"
	"hytx_manager/models"
	"hytx_manager/pkg/setting"
	"strconv"
	"time"
)

/**
	互推圈投诉列表
 */
func InterPushComplaintList(c *gin.Context) {

	item := c.Query("item")
	criminal := c.Query("criminal")
	status := c.Query("status")
	beginTime := c.Query("begin_time")
	endTime := c.Query("end_time")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))
	result, paginate := models.GetInterPushReport(item, status, criminal, beginTime, endTime, page, limit)

	render(c, gin.H{
		"paginate": paginate,
		"data":     result,
	})
}

/**
	互推圈投诉详情
 */
func InterPushComplaintGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := models.InterPushReportGet(id)
	render(c, data)
}

/**
	互推圈投诉处理
 */
func InterPushComplaintSolve(c *gin.Context) {
	id := c.PostForm("id")
	item := c.PostForm("item")
	replay := c.PostForm("replay")
	user := auth.User(c)
	in := models.GetInterPushById(id)
	var operation string
	var operationTime int
	switch item {
	case "1":
		//封停一天
		operation = "封停一天"
		operationTime = 1
	case "3":
		operation = "封停三天"
		operationTime = 3

	}

	models.DB.Table("users").Where("id=?", in.MatchId).Updates(map[string]interface{}{"is_enabled": 1, "enabled_at": time.Now().AddDate(0, 0, operationTime)})
	models.DB.Table("interpush_report").Where("id=?", in.ID).Updates(map[string]interface{}{"status": 1, "solve_user_id": user.ID, "operation": operation, "replay": replay})

	success(c)
}
func InterPushComplaintDel(c *gin.Context) {
	id := c.Query("id")
	models.DB.Delete(&models.InterPushReport{}, "id=?", id)
	success(c)
}
/**
	意见反馈
 */
func Feedback(c *gin.Context) {
	status := c.Query("status")
	beginTime := c.Query("begin_time")
	endTime := c.Query("end_time")
	phone := c.Query("phone")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))

	result, paginate := models.FeedbackList(status, beginTime, endTime, phone, page, limit)

	render(c, gin.H{
		"paginate": paginate,
		"data":     result,
	})
}

/**
	意见反馈详情
 */
func FeedbackGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := models.FeedbackGet(id)

	render(c, data)
}

/**
	意见反馈反馈
 */
func FeedbackSolve(c *gin.Context) {
	id := c.PostForm("id")
	item := c.PostForm("item")
	replay := c.PostForm("replay")
	user := auth.User(c)

	models.DB.Table("feedback").Where("id=?", id).Updates(map[string]interface{}{"solve_user_id": user.ID, "status": 1, "operation": item, "replay": replay})
	success(c)
}

func FeedbackDel(c *gin.Context) {
	id := c.Query("id")
	models.DB.Delete(&models.Feedback{}, "id=?", id)
	success(c)
}
/**
	评论举报列表
 */
func CommentComplaintList(c *gin.Context) {
	item := c.Query("item")
	criminal := c.Query("criminal")
	status := c.Query("status")
	beginTime := c.Query("begin_time")
	endTime := c.Query("end_time")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))

	result, paginate := models.GetCommentFeedbackList(item, status, criminal, beginTime, endTime, page, limit)

	render(c, gin.H{
		"paginate": paginate,
		"data":     result,
	})
}

/**
	评论举报详情
 */
func CommentComplaintGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := models.CommentFeedbackGet(id)
	render(c, data)
}

/**
	评论举报反馈
 */
func CommentComplaintSolve(c *gin.Context) {
	id := c.PostForm("id")
	item := c.PostForm("item")
	replay := c.PostForm("replay")
	user := auth.User(c)

	models.DB.Table("comment_feedback").Where("id=?", id).Updates(map[string]interface{}{"solve_user_id": user.ID, "status": 1, "operation": item, "replay": replay})
	success(c)
}
func CommentComplaintDel(c *gin.Context) {
	id := c.Query("id")
	models.DB.Delete(&models.CommentFeedback{}, "id=?", id)
	success(c)
}
/**
	文章举报列表
 */
func ArticleComplaintList(c *gin.Context) {
	item := c.Query("item")
	criminal := c.Query("criminal")
	status := c.Query("status")
	beginTime := c.Query("begin_time")
	endTime := c.Query("end_time")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))
	result, paginate := models.ArticleComplaintList(page, limit, item, status, criminal, beginTime, endTime)

	render(c, gin.H{
		"paginate": paginate,
		"data":     result,
	})
}
/**
	文章举报详情
 */
func ArticleComplaintGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := models.ComplaintOne(id, 2)
	render(c, data)
}
/**
	文章举报反馈
 */
func ArticleComplaintSolve(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	item := c.PostForm("item")
	replay := c.PostForm("replay")
	user := auth.User(c)
	models.DB.Table("complaint").Where("id=?",id).Updates(map[string]interface{}{"solve_user_id": user.ID, "status": 1, "operation": item, "replay": replay})
	co := models.GetComplaintById(id)
	if item == "1" {
		//下线
		models.DB.Table("articles").Where("id=?",co.Quote).Update("status", 6)
	}
	if item == "2" {
		//下线
		models.DB.Table("articles").Where("id=?",co.Quote).Update("status", 6)
		models.DB.Table("media_users").Where("id=?", co.TargetId).Update("is_enabled", 1)
	}
	success(c)
}
/**
	网页举报列表
 */
func PageComplaintList(c *gin.Context) {
	item := c.Query("item")
	criminal := c.Query("criminal")
	status := c.Query("status")
	beginTime := c.Query("begin_time")
	endTime := c.Query("end_time")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(setting.AppSetting.PageSize)))
	result, paginate := models.PageComplaintList(page, limit, item, status, criminal, beginTime, endTime)

	render(c, gin.H{
		"paginate": paginate,
		"data":     result,
	})
}
/**
	网页举报详情
 */
func PageComplaintGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := models.ComplaintOne(id, 4)
	render(c, data)
}
/**
	网页举报反馈
 */
func PageComplaintSolve(c *gin.Context) {
	id := c.PostForm("id")
	item := c.PostForm("item")
	replay := c.PostForm("replay")
	user := auth.User(c)
	models.DB.Table("complaint").Where("id=?",id).Updates(map[string]interface{}{"solve_user_id": user.ID, "status": 1, "operation": item, "replay": replay})
	success(c)
}
func ComplaintDel(c *gin.Context) {
	id := c.Query("id")
	models.DB.Delete(&models.Complaint{}, "id=?", id)
	success(c)
}