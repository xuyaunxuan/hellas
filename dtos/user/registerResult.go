package user

type RegisterResult struct {
	AccountId string `json:"id" binding:"required,alphanum,max=50"`
	MailAddress string `json:"mail" binding:"required,email,max=50"`
	NickName string `json:"nickName" binding:"required,max=20"`
	Password string `json:"password" binding:"required,alphanum,max=20"`
}