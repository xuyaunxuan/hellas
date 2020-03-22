package routers

import (
	"github.com/gin-gonic/gin"
	"hellas/common/setting"
	"hellas/common/utils"
	"hellas/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	user := r.Group("/user")
	user.Use(utils.AuthorityCheck())
	{
		// 新用户注册
		user.POST("/register", controller.RegisterUser)
		// 请求邮箱验证码
		user.POST("/sendCaptchaMail", controller.SendUserCaptchaMail)
		// 密码变更
		user.POST("/resetPassword", controller.ResetUserPassword)
		// 用户登录
		user.POST("/login", controller.Login)
		// 用户信息编辑
		user.POST("/editDetail", controller.EditUserDetail)
	}

	b := r.Group("/b")
	b.Use(utils.AuthorityCheck())
	{
		// 投稿
		b.POST("/subscribe", controller.Subscribe)
		// 获取用户投稿
		b.POST("/myArticles", controller.ViewUserArticle)
		// 编辑投稿
		b.POST("/editArticles", controller.EditSubscribe)
		// 删除投稿
		b.POST("/deleteArticles", controller.DeleteSubscribe)
		// 获取最新文章
		b.POST("/showArticles", controller.ViewArticle)
		// 获取文章详情
		b.GET("/articleDetail", controller.ArticleDetail)
	}

	return r
}