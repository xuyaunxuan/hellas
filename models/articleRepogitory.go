package models

import (
	"hellas/common/constant"
	"hellas/common/utils"
	"hellas/dtos"
	"hellas/dtos/article"
	"hellas/dtos/common"
	"strings"
	"time"
)

// 投稿新文章
func CreateNewArticle(param article.SubscribeParameter, accountId string) common.BaseResult {
	// 开启事务
	tx := db.Begin()
	var result common.BaseResult
	var count int
	var userPostCount int
	// 当前最大文章数检索
	db.Model(&dtos.Article{}).Count(&count)
	// 当前用户最大文章数检索
	db.Model(&dtos.Article{}).Where(&dtos.Article{AccountId: accountId}).Count(&userPostCount)
	// 默认公开
	var privateFlg = "0"
	if param.IsPrivate {
		privateFlg = "1"
	}
	var article = dtos.Article{
		Id: utils.CreateMaxPostPath(count),
		AccountId: accountId,
		Sequence: userPostCount + 1,
		Title: param.Title,
		ContentOrigin: param.ArticleOri,
		Content: param.Article,
		PrivateFlg: privateFlg,
		DeleteFlg : "0",
		InsertDateTime: time.Now(),
		UpdateDateTime: time.Now(),
		Tag: param.Tag,
	}

	// 插入文章
	db.Create(&article)
	// 提交事务
	tx.Commit()
	// 处理成功
	result.Result = constant.OK
	return result
}

// 编辑投稿文章
func EditArticle(param article.SubscribeParameter, accountId string) common.BaseResult {
	var result common.BaseResult
	if param.ArticlePath == "" {
		// 该文章不存在
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("","articleNotExist"))
		return result
	}
	// 开启事务
	tx := db.Begin()
	var count int
	// 文章有效性检索
	db.Model(dtos.Article{}).Where(&dtos.Article{Id:param.ArticlePath,AccountId:accountId,DeleteFlg:"0"}).Count(&count)
	if count == 0 {
		// 该文章不存在
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("","articleNotExist"))
		return result
	}

	// 默认公开
	var privateFlg = "0"
	if param.IsPrivate {
		privateFlg = "1"
	}
	// 更新文章
	db.Model(dtos.Article{}).Where(dtos.Article{Id:param.ArticlePath}).Update(dtos.Article{Title : param.Title, ContentOrigin:param.ArticleOri,Content:param.Article,PrivateFlg:privateFlg,Tag:param.Tag, UpdateDateTime:time.Now()})

	// 提交事务
	tx.Commit()
	// 处理成功
	result.Result = constant.OK
	return result
}

func DeleteArticle(param article.DeleteParameter, accountId string) common.BaseResult {
	var result common.BaseResult
	if param.ArticlePath == "" {
		// 该文章不存在
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("","articleNotExist"))
		return result
	}
	// 开启事务
	tx := db.Begin()

	var count int
	// 文章有效性检索
	db.Model(dtos.Article{}).Where(&dtos.Article{Id:param.ArticlePath,AccountId:accountId,DeleteFlg:"0"}).Count(&count)
	if count == 0 {
		// 该文章不存在
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("","articleNotExist"))
		return result
	}

	// 更新文章
	db.Model(dtos.Article{}).Where(dtos.Article{Id:param.ArticlePath}).Update(dtos.Article{DeleteFlg:"1", UpdateDateTime:time.Now()})

	// 提交事务
	tx.Commit()
	// 处理成功
	result.Result = constant.OK
	return result
}

// 按投稿时间降顺获取用户文章
func GetUserArticle(param article.ViewArticleParameter, accountId string) article.ViewArticleResult {
	var result article.ViewArticleResult
	var articles []article.ViewArticle
	// 文章信息DTO
	var articleList []dtos.Article
	db.Where(&dtos.Article{AccountId: accountId,DeleteFlg:"0"}).Offset(param.Offset).Limit(param.Limit).Order("id desc").Find(&articleList)
	for _, art := range articleList {
		var rs = article.ViewArticle {
			Title: art.Title,
			ArticlePath: art.Id,
			AccountId:art.AccountId,
			PostTime:art.InsertDateTime,
			Tag:art.Tag,
			IsPrivate:art.PrivateFlg,
		}
		articles = append(articles, rs)
	}
	result.Articles = articles
	result.Result = constant.OK
	return result
}

// 按投稿时间降顺获取公开的文章
func GetArticle(param article.ViewArticleParameter) article.ViewArticleResult {
	var result article.ViewArticleResult
	var articles []article.ViewArticle
	// 文章信息DTO
	var articleList []dtos.Article
	// 按投稿时间降顺获取公开的文章
	db.Where(&dtos.Article{PrivateFlg: "0",DeleteFlg:"0"}).Offset(param.Offset).Limit(param.Limit).Order("id desc").Find(&articleList)
	for _, art := range articleList {
		var rs = article.ViewArticle {
			Title: art.Title,
			ArticlePath: art.Id,
			AccountId:art.AccountId,
			PostTime:art.InsertDateTime,
			Tag:art.Tag,
		}
		articles = append(articles, rs)
	}
	result.Articles = articles
	result.Result = constant.OK
	return result
}

//文章详细内容
func GetArticleDetail(postId string, remoteAddr string) article.ViewArticleResult {
	var result article.ViewArticleResult
	var articles []article.ViewArticle
	// 文章信息DTO
	var articleList []dtos.Article
	db.Where(&dtos.Article{Id: postId}).Find(&articleList)
	for _, art := range articleList {
		var rs = article.ViewArticle {
			Title: art.Title,
			ArticlePath: art.Id,
			AccountId:art.AccountId,
			PostTime:art.InsertDateTime,
			UpdateTime:art.UpdateDateTime,
			Tag:art.Tag,
			Content:art.Content,
			ContentOri:art.ContentOrigin,
			IsPrivate:art.PrivateFlg,
		}
		articles = append(articles, rs)
	}
	result.Articles = articles
	if len(articles) == 0 {
		result.Result = constant.NG
	} else {
		var addr = strings.Split(remoteAddr, ":")
		// 访问信息
		var articleVisits []dtos.ArticleVisits
		db.Where(&dtos.ArticleVisits{Id: postId,ViewAddr:addr[0]}).Find(&articleVisits)
		if len(articleVisits) == 0 {
			var articleVisit = dtos.ArticleVisits{
				Id: postId,
				ViewAddr:addr[0],
				ViewTime:time.Now(),
			}
			// 插入文章
			db.Create(&articleVisit)
		} else {
			db.Model(dtos.ArticleVisits{}).Where(dtos.ArticleVisits{Id:postId,ViewAddr:addr[0]}).Update(dtos.ArticleVisits{ViewTime:time.Now()})
		}
		result.Result = constant.OK
	}

	return result
}