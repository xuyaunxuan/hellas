package user

type ResetPasswordParameter struct {
	MailAddress string `binding:"required,email,max=50"`
	CaptchaCode string `binding:"required,alphanum,len=6"`
	OncePassword string `binding:"required,alphanum,max=20"`
	TwicePassword string `binding:"required,alphanum,max=20"`
}
