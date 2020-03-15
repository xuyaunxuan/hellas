package dtos

import "time"

// 对应数据库User表
type User struct {
	Id int
	AccountId string
	MailAddress string
	NickName string
	Password string
	Salt string
	InsertDateTime time.Time
	UpdateDateTime time.Time
}