package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"hellas/common/constant"
	"hellas/common/e"
	"hellas/common/utils"
	"hellas/dtos/common"
	"hellas/dtos/user"
	"hellas/models"
	"log"
	"net/http"
)

// 新用户注册
func RegisterUser(c *gin.Context) {
	var registerParameter user.RegisterParameter
	//c.Query("username")
	// 参数验证
	if err := c.ShouldBindJSON(&registerParameter); err != nil {
		log.Printf("%+v", registerParameter)
		var errorDto common.ErrorDto
		errorDto.Message = e.MsgFlags[e.INVALID_PARAMS] +  err.Error()
		// 生成错误信息
		errorDto.Errors = utils.CreateMessages(err.(validator.ValidationErrors))
		// 返回status400
		c.JSON(http.StatusBadRequest, errorDto)
		return
	}

	log.Printf("%+v", registerParameter)
	// 新用户创建
	result := models.CreateNewUser(registerParameter)
	if result.Result == constant.NG {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}