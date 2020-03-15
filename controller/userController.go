package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"hellas/common/constant"
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
	// 参数验证
	if err := c.ShouldBindJSON(&registerParameter); err != nil {
		log.Printf("%+v", registerParameter)
		var errorDto common.ErrorDto
		// 生成错误信息
		errorDto.Errors = utils.CreateMessages(err.(validator.ValidationErrors))
		// 返回status400
		c.JSON(http.StatusBadRequest, errorDto)
		return
	}

	// 新用户创建
	result := models.CreateNewUser(registerParameter)
	if result.Result == constant.NG {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

// 发送用户验证邮件
func SendUserCaptchaMail(c *gin.Context) {
	var sendCaptchaMailParameter user.SendCaptchaMailParameter
	// 参数验证
	if err := c.ShouldBindJSON(&sendCaptchaMailParameter); err != nil {
		var errorDto common.ErrorDto
		// 生成错误信息
		errorDto.Errors = utils.CreateMessages(err.(validator.ValidationErrors))
		// 返回status400
		c.JSON(http.StatusBadRequest, errorDto)
		return
	}

	// 发送验证邮件
	result := models.SendCaptchaMail(sendCaptchaMailParameter)
	if result.Result == constant.NG {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

// 修改用户密码
func ResetUserPassword(c *gin.Context) {
	var resetPasswordParameter user.ResetPasswordParameter
	// 参数验证
	if err := c.ShouldBindJSON(&resetPasswordParameter); err != nil {
		var errorDto common.ErrorDto
		// 生成错误信息
		errorDto.Errors = utils.CreateMessages(err.(validator.ValidationErrors))
		// 返回status400
		c.JSON(http.StatusBadRequest, errorDto)
		return
	}

	// 重置密码
	result := models.ResetPassword(resetPasswordParameter)
	if result.Result == constant.NG {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

// 用户登录
func Login(c *gin.Context) {
	var loginParameter user.LoginParameter
	// 参数验证
	if err := c.ShouldBindJSON(&loginParameter); err != nil {
		var errorDto common.ErrorDto
		// 生成错误信息
		errorDto.Errors = utils.CreateMessages(err.(validator.ValidationErrors))
		// 返回status400
		c.JSON(http.StatusBadRequest, errorDto)
		return
	}

	// 用户登录
	result := models.Login(loginParameter)
	if result.Result == constant.NG {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}