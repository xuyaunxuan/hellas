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

		user := r.Group("/v1/u")
	user.Use(utils.AuthorityCheck())
	{
		// 新用户注册
		user.POST("/reg", controller.RegisterUser)
		// 请求邮箱验证码
		user.POST("/send/mail", controller.SendUserCaptchaMail)
		// 密码变更
		user.POST("/reset/password", controller.ResetUserPassword)
		// 用户登录
		user.POST("/login", controller.Login)
		// 用户信息编辑
		//user.POST("/edit/info", controller.EditUserDetail)
	}

	b := r.Group("/v1/b")
	b.Use(utils.AuthorityCheck())
	{
		// 投稿
		b.POST("/u/post", controller.Subscribe)
		// 获取用户投稿
		b.POST("/u/all/post", controller.ViewUserArticle)
		// 编辑投稿
		b.POST("/u/edit", controller.EditSubscribe)
		// 删除投稿
		b.POST("/u/del/post", controller.DeleteSubscribe)
		// 获取最新文章
		b.POST("/all/post", controller.ViewArticle)
		// 获取文章详情
		b.GET("/detail", controller.ArticleDetail)
	}

	return r
}