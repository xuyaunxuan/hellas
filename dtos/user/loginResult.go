package user

import "hellas/dtos/common"

type LoginResult struct {
	common.BaseResult
	AccountId string `json:"accountId"`
	MailAddress string `json:"mailAddress"`
	NickName string `json:"nickName"`
	Token string `json:"token""`
}
