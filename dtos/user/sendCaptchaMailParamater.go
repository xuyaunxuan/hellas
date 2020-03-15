package user


type SendCaptchaMailParameter struct {
	MailAddress string `binding:"required,email,max=50"`
}

