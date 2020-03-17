package routers

import (
	"github.com/gin-gonic/gin"
	"hellas/common/utils"
	"hellas/controller"

	"hellas/common/setting"
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

	return r
}