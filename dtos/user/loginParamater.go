package user

type LoginParameter struct {
	AccountId string `binding:"max=50"`
	MailAddress string `binding:"max=50"`
	Password string `binding:"required,alphanum,max=20"`
}
