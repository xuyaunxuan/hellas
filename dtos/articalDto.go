package dtos

import "time"

type Article struct {
	AccountId string
	Sequence int
	Title string
	Content string
	PrivateFlg string
	DeleteFlg string
	InsertDateTime time.Time
	UpdateDateTime time.Time
	Tag string
}

