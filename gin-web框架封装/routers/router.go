package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"hytx_manager/controller"
	"hytx_manager/middleware/auth"
	"hytx_manager/middleware/cros"
	"hytx_manager/pkg/setting"

	"github.com/gin-contrib/gzip"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// gzip 压缩
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.Use(gin.Logger())
	r.Use(handleErrors())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	// 跨域
	r.Use(cros.Cors())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
		return
	})
	//root组下接口
	root := r.Group("/api")

	root.Use(auth.RegisterGlobalAuth())
	// 登陆
	root.POST("/login", controller.Login)
	rbac := root.Group("/")
	rbac.Use(auth.AuthMiddleware())
	rbac.Use(auth.PermissionCheck())
	{
		//角色
		rbac.GET("/get_role", controller.GetRoles)
		rbac.GET("/get_role_no_page", controller.GetRolesNoPage)
		rbac.POST("/post_role", controller.AddRole)
		rbac.GET("/get_role_id", controller.GetRole)
		rbac.POST("/update_role_id", controller.UpdateRole)
		rbac.POST("/delete_role_id", controller.DeletedRole)
		rbac.POST("/get_role_is_enabled", controller.IsEnabledRole)
		// 角色对应的权限
		rbac.GET("/get_role_permission", controller.GetRolePermissions)
		rbac.POST("/post_role_permission", controller.UpdateRolePermissions)

		// 权限
		rbac.GET("/get_permission", controller.GetPermissions)
		rbac.GET("/get_permission_id", controller.GetPermissionById)
		rbac.POST("/post_permission", controller.AddPermission)
		rbac.POST("/update_permission", controller.UpdatePermission)
		rbac.POST("/deleted_permission", controller.DeletedPermission)

		// permissionMenu
		rbac.GET("/get_permission_menu", controller.GetPermissionMenus)

		// admin
		rbac.GET("/admin_role_permission", controller.GetAdminRolePermission)
		rbac.GET("/get_admins", controller.GetAdmins)
		rbac.GET("/get_admin_id", controller.GetAdminById)
		rbac.POST("/post_admin", controller.AddAdmin)
		rbac.POST("/update_admin", controller.UpdateAdmin)
		rbac.POST("/deleted_admin", controller.DeletedAdmin)
		rbac.POST("/able_admin", controller.AdminAble)
	}
	au := root.Group("/")
	au.Use(auth.AuthMiddleware())
	au.Use(auth.PermissionCheck())
	//自媒体
	media := au.Group("/media")
	{
		// 账户审核列表
		media.GET("/get_users", controller.GetMediaUsers)
		media.GET("/get_user_by_id", controller.GetMediaUserInfoById)
		media.GET("/update_user_status", controller.UpdateMediaUserStatus)
		media.GET("/update_user_is_enabled", controller.MediaAble)
		article := media.Group("/article")
		{
			// 文章标签
			article.GET("/get_labels", controller.GetArticleLabel)
			article.GET("/get_label_by_id", controller.GetArticleLabelById)
			article.POST("/post_label", controller.AddArticleLabel)
			article.GET("/update_label", controller.UpdateArticleLabel)
			article.GET("/deleted_article_label", controller.DeletedArticleLabel)
			article.GET("/update_ad_id", controller.UpdateArticleAdvertisementId)
			// 文章分类
			article.GET("/get_categories", controller.GetArticleCategories)
			article.GET("/get_category_by_id", controller.GetArticleCategoryById)
			article.POST("/post_category", controller.AddArticleCategory)
			article.GET("/update_category", controller.UpdateArticleCategory)
			article.GET("/deleted_category", controller.DeletedArticleCategory)
			article.GET("/suggest_article_category", controller.SuggestArticleCategoryList)
			article.GET("/suggest_article_count", controller.SuggestArticleCategoryCount)
			// 文章管理
			article.GET("/get_articles", controller.GetArticles)
			article.GET("/get_article_by_id", controller.GetArticleById)
			article.GET("/update_article", controller.UpdateArticle)
		}
	}
	//会员列表
	member := au.Group("/member")
	{
		member.GET("/get_member", controller.GetMembers)
		member.GET("/get_member_info_by_id", controller.GetMemberInfoById)
		member.GET("/get_member_template_by_id", controller.GetPersonalPageTemplateUser)
		member.GET("/get_member_marketing_template_by_id", controller.GetMemberMarketing)
		member.POST("/able", controller.MemberAble)
		// 会员充值（需涉及分销提成）
		member.POST("/recharge", controller.MemberRecharge)
	}

	//运营商列表
	operators := au.Group("/operators")
	{
		operators.GET("/list", controller.OperatorsList)
		operators.GET("/get", controller.OperatorsGet)
		operators.POST("/add", controller.OperatorAdd)
		operators.PUT("/edit", controller.OperatorEdit)
		operators.DELETE("/del", controller.OperatorDel)
		operators.POST("/able", controller.OperatorAble)
		operators.GET("/manage_center_list", controller.OperatorManageCenterList)
		operators.POST("/examine", controller.OperatorExamine)
	}
	// 城市合伙人
	partner := au.Group("/partner")
	{
		partner.GET("/list", controller.PartnerList)
		partner.GET("/get", controller.PartnerGet)
		partner.POST("/add", controller.PartnerAdd)
		partner.POST("/edit", controller.PartnerEdit)
		partner.GET("/del", controller.PartnerDel)
		partner.GET("/able", controller.PartnerAble)
		partner.GET("/operator_list", controller.PartnerOperatorList)
		partner.GET("/is_enabled", controller.IsEnabledPartner)
		partner.GET("/examine_partner", controller.ExaminePartner)
		partner.GET("/down_grade_partner", controller.DowngradePartner)
		partner.GET("/get_operator_list", controller.GetOperatorList)
		partner.GET("/get_member_list", controller.GetMemberList)
	}

	// vip_fee_config
	root.GET("/vip_fee_config", controller.GetVipFeeConfigs)

	root.GET("/index", controller.Index)
	//广告
	ad := au.Group("/ad")
	{
		ad.GET("/list", controller.AdList)
		ad.GET("/get", controller.AdGet)
		ad.POST("/add", controller.AddAd)
		ad.PUT("/edit", controller.EditAd)
		ad.DELETE("/del", controller.DElAd)
		ad.POST("/ban", controller.AdBan)
	}
	//投诉管理
	complaint := au.Group("/complaint")
	{
		complaint.GET("/inter_push", controller.InterPushComplaintList)
		complaint.GET("/inter_push/get", controller.InterPushComplaintGet)
		complaint.POST("/inter_push/solve", controller.InterPushComplaintSolve)
		complaint.DELETE("/inter_push/del", controller.InterPushComplaintDel)
		complaint.GET("/article", controller.ArticleComplaintList)
		complaint.GET("/article/get", controller.ArticleComplaintGet)
		complaint.POST("/article/solve", controller.ArticleComplaintSolve)
		complaint.DELETE("/article/del", controller.ComplaintDel)
		complaint.GET("/page", controller.PageComplaintList)
		complaint.GET("/page/get", controller.PageComplaintGet)
		complaint.POST("/page/solve", controller.PageComplaintSolve)
		complaint.DELETE("/page/del", controller.ComplaintDel)
		complaint.GET("/feedback", controller.Feedback)
		complaint.GET("/feedback/get", controller.FeedbackGet)
		complaint.POST("/feedback/solve", controller.FeedbackSolve)
		complaint.DELETE("/feedback/del", controller.FeedbackDel)
		complaint.GET("/comment", controller.CommentComplaintList)
		complaint.GET("/comment/get", controller.CommentComplaintGet)
		complaint.POST("/comment/solve", controller.CommentComplaintSolve)
		complaint.DELETE("/comment/del", controller.CommentComplaintDel)
	}
	financial := au.Group("/financial")
	{
		financial.GET("/put_forward", controller.PutForward)
		financial.POST("/put_forward/make", controller.MakeMoney)
		financial.GET("/user_recharges", controller.UserRechargesList)
		financial.GET("/user_profit", controller.UserProfitList)
		financial.GET("/user_profit/detail", controller.UserProfitDetail)
		financial.GET("/user_profit/put", controller.PutForwardList)
		financial.GET("/distributor_forward", controller.PutForwardListByDistributor)
		financial.POST("/distributor_forward/put", controller.PutForwardByDistributor)
		financial.GET("/distributor", controller.DistributorProfitLog)
		financial.GET("/distributor/detail", controller.DistributorProfitLogList)
		financial.GET("/maibao/list", controller.MaibaoLog)
		financial.GET("/maibao/recharge_stream", controller.MaibaoLogByUser)
		financial.POST("/maibao/recharge", controller.MaibaoRecharge)
		financial.GET("/maibao/spending", controller.MaibaoSpendingLog)
		financial.GET("/maibao/spending/detail", controller.MaibaoGain)
		financial.GET("/manager_center/recharges", controller.ManagerCenterRecharges)
		financial.GET("/manager_center/recharges_detail", controller.ManagerCenterRechargesDetail)
		financial.GET("/manager_center/recharges_cash", controller.ManagerCenterRechargesCash)
		financial.POST("/manager_center/recharges_cash_status", controller.ManagerCenterRechargesCashStatus)
		financial.GET("/operator/recharges", controller.OperatorRecharges)
		financial.GET("/operator/recharges_detail", controller.OperatorRechargesDetail)
		financial.GET("/operator/recharges_cash", controller.OperatorRechargesCash)
		financial.POST("/operator/recharges_cash_status", controller.OperatorRechargesCashStatus)
		financial.GET("/partner/recharges", controller.PartnerRecharges)
		financial.GET("/partner/recharges_detail", controller.PartnerRechargesDetail)
		financial.GET("/partner/recharges_cash", controller.PartnerCenterRechargesCash)
		financial.POST("/partner/recharges_cash_status", controller.PartnerCenterRechargesCashStatus)
	}
	root.GET("/operation_log", controller.OperationLogList)
	account := au.Group("/account").Use(auth.AuthMiddleware())
	{
		account.GET("/profile", controller.Profile)
	}

	// 管理中心
	manage := au.Group("manage_center")
	{
		manage.GET("/list", controller.ManageCenterList)
		manage.POST("/add", controller.ManageCenterAdd)
		manage.PUT("/edit", controller.ManageCenterEdit)
		manage.GET("/get", controller.ManageCenterGet)
		manage.DELETE("/del", controller.ManageCenterDel)
		manage.POST("/able", controller.ManageCenterAble)

	}
	pageTemplate := au.Group("page_template")
	{
		pageTemplate.GET("fixed_template", controller.FixedTemplateList)
		pageTemplate.GET("fixed_template/get", controller.FixedTemplateGet)
		pageTemplate.POST("fixed_template/add", controller.AddFixedTemplate)
		pageTemplate.PUT("fixed_template/edit", controller.EditFixedTemplate)
		pageTemplate.DELETE("fixed_template/del", controller.DelFixedTemplate)
		pageTemplate.POST("fixed_template/able", controller.AbleFixedTemplate)

		pageTemplate.GET("fixed_template_category", controller.FixedTemplateCategoryList)
		pageTemplate.GET("fixed_template_category/get", controller.FixedTemplateCategoryGet)
		pageTemplate.POST("fixed_template_category/add", controller.AddFixedTemplateCategory)
		pageTemplate.PUT("fixed_template_category/edit", controller.EditFixedTemplateCategory)
		pageTemplate.DELETE("fixed_template_category/del", controller.DelFixedTemplateCategory)
		pageTemplate.POST("fixed_template_category/able", controller.AbleFixedTemplateCategory)

		pageTemplate.GET("custom_category", controller.CustomCategoryList)
		pageTemplate.GET("custom_category/get", controller.CustomCategoryGet)
		pageTemplate.POST("custom_category/add", controller.AddCustomCategory)
		pageTemplate.PUT("custom_category/edit", controller.EditCustomCategory)
		pageTemplate.DELETE("custom_category/del", controller.DelCustomCategory)
		pageTemplate.POST("custom_category/able", controller.AbleCustomCategory)

		pageTemplate.GET("custom_component", controller.CustomComponentList)
		pageTemplate.GET("custom_component/get", controller.CustomComponentGet)
		pageTemplate.POST("custom_component/add", controller.AddCustomComponent)
		pageTemplate.PUT("custom_component/edit", controller.EditCustomComponent)
		pageTemplate.DELETE("custom_component/del", controller.DelCustomComponent)
		pageTemplate.POST("custom_component/able", controller.AbleCustomComponent)

		pageTemplate.GET("custom_background", controller.CustomBackgroundList)
		pageTemplate.GET("custom_background/get", controller.CustomBackgroundGet)
		pageTemplate.POST("custom_background/add", controller.AddCustomBackground)
		pageTemplate.PUT("custom_background/edit", controller.EditCustomBackground)
		pageTemplate.DELETE("custom_background/del", controller.DelCustomBackground)
		pageTemplate.POST("custom_background/able", controller.AbleCustomBackground)
	}

	// 全名推广
	spread := au.Group("spread")
	{
		spread.GET("/list", controller.GetSpreads)
		spread.GET("/get_by_id", controller.GetSpreadById)
		spread.DELETE("/del", controller.DelSpread)
		spread.PUT("/update", controller.UpdateSpread)
	}
	//经销商
	distributor := au.Group("distributor")
	{
		distributor.GET("/list", controller.DistributorList)
		distributor.GET("/get", controller.DistributorGet)
		distributor.POST("/add", controller.DistributorAdd)
		distributor.POST("/edit", controller.DistributorEdit)
		distributor.POST("/disable", controller.DistributorDisable)
		distributor.GET("/re_log", controller.DIstributorLog)
	}

	// 企业
	enterprise := au.Group("/enterprise")
	{
		enterprise.GET("/list", controller.EnterpriseList)
		enterprise.GET("/get-staff-list/:enterprise_id", controller.GetStaffList)
		enterprise.POST("/account-disable/:uid", controller.DisableAccount)
		enterprise.GET("/ident-list", controller.IdentList)
		enterprise.GET("/get/:id", controller.EnterInfo)
		enterprise.POST("/ident/:id", controller.Ident)
		enterprise.GET("/staff-recharges", controller.StaffRecharges)
		enterprise.GET("/self-list", controller.EnterpriseSelfList)
		enterprise.GET("/enterprise-staff-count", controller.StaffCount)
		enterprise.GET("/enterprise-user-count", controller.UserCount)
		enterprise.GET("/enterprise-recharges", controller.EnterpriseRecharges)
		enterprise.POST("/enterprise-vip", controller.EnterpriseVip)
		enterprise.POST("/add-enterad", controller.AddEnterAd)
		enterprise.GET("/get-enterads", controller.GetEnterAds)
		enterprise.GET("/get-enterad/:id", controller.GetEnterAd)
		enterprise.POST("/edit-enterad/:id", controller.EditEnterAd)
		enterprise.GET("/del-enterad/:id", controller.DeleteEnterAd)
		enterprise.POST("/disable-enterad/:id", controller.DisableEnterAd)
		enterprise.GET("/get-resers", controller.GetResers)
		enterprise.GET("/use-reser/:id", controller.UseReserCode)
		enterprise.GET("/enterprise-resers", controller.EnterResers)
		enterprise.POST("/add-vipcode", controller.AddVipCode)
		enterprise.GET("/get-vipcodes", controller.GetVipCodeList)
	}
	//app广告
	app := au.Group("/app")
	{
		app.POST("/add-push-msg", controller.AddPushMsg)
		app.GET("/get-push-msg/:id", controller.GetPushMsg)
		app.GET("/get-push-msgs", controller.GetPushMsgs)
		app.DELETE("/del-push-msg/:id", controller.DeletePushMsg)
		app.POST("/edit-push-msg/:id", controller.EditPushMsg)

		app.POST("/add-banner", controller.AddBanner)
		app.GET("/get-banner/:id", controller.GetBanner)
		app.GET("/get-banners", controller.GetBanners)
		app.DELETE("/del-banner/:id", controller.DeleteBanner)
		app.POST("/edit-banner/:id", controller.EditBanner)
	}

	root.POST("/send_sms_code", controller.SendSMSCode)
	root.GET("/cities", controller.GetCity)
	au.GET("/test", controller.AuthTest)
	return r
}
