package article

import (
	"hellas/dtos/common"
)

type ViewArticleResult struct {
	common.BaseResult
	Articles []ViewArticle `json:"articles"`
}