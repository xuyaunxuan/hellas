package models

import (
	"hellas/common/constant"
	"hellas/dtos"
	"hellas/dtos/article"
	"hellas/dtos/common"
	"time"
)

// 投稿新文章
func CreateNewArticle(param article.SubscribeParameter, accountId string) common.BaseResult {
	// 开启事务
	tx := db.Begin()
	var result common.BaseResult
	var count int
	// 当前用户最大文章数检索
	db.Model(&dtos.Article{}).Where(&dtos.Article{AccountId: accountId}).Count(&count)
	// 默认公开
	var privateFlg = "0"
	if param.IsPrivate {
		privateFlg = "1"
	}
	var article = dtos.Article{
		AccountId: accountId,
		Sequence: count+ 1,
		Title: param.Title,
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
