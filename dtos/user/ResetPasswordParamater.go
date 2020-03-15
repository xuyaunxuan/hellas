package user

type ResetPasswordParameter struct {
	MailAddress string `json:"mail" binding:"required,email,max=50"`
	CaptchaCode string `json:"captchaCode" binding:"required,alphanum,len=6"`
	OncePassword string `json:"oncePassword" binding:"required,alphanum,max=20"`
	TwicePassword string `json:"twicePassword" binding:"required,alphanum,max=20"`
}
