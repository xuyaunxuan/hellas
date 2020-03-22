package article

import (
	"time"
)

type ViewArticle struct {
	Title string `json:"title"`
	ArticlePath string `json:"articlePath"`
	AccountId string `json:"accountId"`
	PostTime time.Time `json:"postTime"`
	UpdateTime time.Time `json:"updateTime"`
	Tag string `json:"tag"`
	Content string `json:"content"`
	IsPrivate string `json:"isPrivate"`
}
