package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"hellas/common/constant"
	"hellas/common/utils"
	"hellas/dtos/article"
	"hellas/dtos/common"
	"hellas/models"
	"net/http"
)

// 获取所有公开文章
func ViewArticle(c *gin.Context) {
	var viewArticleParameter article.ViewArticleParameter
	if err := c.ShouldBindJSON(&viewArticleParameter); err != nil {
		var baseResult common.BaseResult
		// 生成错误信息
		baseResult.ErrorDto.Errors = utils.CreateMessages(err.(validator.ValidationErrors))
		// 返回status400
		c.JSON(http.StatusBadRequest, baseResult)
		return
	}

	result := models.GetArticle(viewArticleParameter)
	c.JSON(http.StatusOK, result)
}

// 打开公开文章
func ArticleDetail(c *gin.Context) {
	postId := c.Query("post")
	result := models.GetArticleDetail(postId, c.Request.RemoteAddr)

	if result.Result == constant.NG {
		// 返回status404
		c.JSON(http.StatusNotFound, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

// 获取用户投稿文章
func ViewUserArticle(c *gin.Context) {
	var viewArticleParameter article.ViewArticleParameter
	if err := c.ShouldBindJSON(&viewArticleParameter); err != nil {
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
	result := models.GetUserArticle(viewArticleParameter, claims.AccountId)
	c.JSON(http.StatusOK, result)
}

// 投稿
func Subscribe(c *gin.Context) {
	var subscribeParameter article.SubscribeParameter
	// 参数验证
	if err := c.ShouldBindJSON(&subscribeParameter); err != nil {
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

// 编辑投稿
func EditSubscribe(c *gin.Context) {
	var subscribeParameter article.SubscribeParameter
	// 参数验证
	if err := c.ShouldBindJSON(&subscribeParameter); err != nil {
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

	result := models.EditArticle(subscribeParameter, claims.AccountId)
	if result.Result == constant.NG {
		c.JSON(http.StatusBadRequest, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

// 删除投稿
func DeleteSubscribe(c *gin.Context) {
	var deleteParameter article.DeleteParameter
	// 参数验证
	if err := c.ShouldBindJSON(&deleteParameter); err != nil {
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

	result := models.DeleteArticle(deleteParameter, claims.AccountId)
	if result.Result == constant.NG {
		c.JSON(http.StatusBadRequest, result)
		return
	}
	c.JSON(http.StatusOK, result)
}
