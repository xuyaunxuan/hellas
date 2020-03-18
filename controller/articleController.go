package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"hellas/common/constant"
	"hellas/common/utils"
	"hellas/dtos/article"
	"hellas/dtos/common"
	"hellas/models"
	"log"
	"net/http"
)

// 投稿
func Subscribe(c *gin.Context) {
	var subscribeParameter article.SubscribeParameter
	// 参数验证
	if err := c.ShouldBindJSON(&subscribeParameter); err != nil {
		log.Printf("%+v", subscribeParameter)
		var baseResult common.BaseResult
		// 生成错误信息
		baseResult.ErrorDto.Errors = utils.CreateMessages(err.(validator.ValidationErrors))
		// 返回status400
		c.JSON(http.StatusBadRequest, baseResult)
		return
	}

	// 从header拿token
	token := c.Request.Header.Get("Authorization")
	// 解析token拿到用户ID
	claims, _ :=utils.JwtParse(token)

	result := models.CreateNewArticle(subscribeParameter, claims.AccountId)
	if result.Result == constant.NG {
		c.JSON(http.StatusBadRequest, result)
		return
	}
	c.JSON(http.StatusOK, result)
}
