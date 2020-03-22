package dtos

import "time"

type Article struct {
	Id string
	AccountId string
	Sequence int
	Title string
	Content string
	ContentOrigin string
	PrivateFlg string
	DeleteFlg string
	InsertDateTime time.Time
	UpdateDateTime time.Time
	Tag string
}

