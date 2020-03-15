package routers

import (
	"github.com/gin-gonic/gin"
	"hellas/controller"

	"ginWork/common/setting"
	"ginWork/routers/api"
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
	}


	adi1 := r.Group("/api/v1")
	{
		//获取标签列表
		adi1.GET("/tags", api.GetTags)
		//新建标签
		adi1.POST("/tags", api.AddTag)
		//更新指定标签
		adi1.PUT("/tags/:id", api.EditTag)
		//删除指定标签
		adi1.DELETE("/tags/:id", api.DeleteTag)
	}

	return r
}