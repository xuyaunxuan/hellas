package routers

import (
	"github.com/gin-gonic/gin"
	"hellas/controller"

	"ginWork/common/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	user := r.Group("/user")
	{
		// 新用户注册
		user.POST("/register", controller.RegisterUser)
		// 请求邮箱验证码
		user.POST("/sendCaptchaMail", controller.SendUserCaptchaMail)
	}

	return r
}