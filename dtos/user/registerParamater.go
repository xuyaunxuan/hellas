package user

type RegisterParameter struct {
	AccountId string `binding:"required,alphanum,max=50"`
	MailAddress string `binding:"required,email,max=50"`
	NickName string `binding:"required,max=20"`
	Password string `binding:"required,alphanum,max=20"`
}