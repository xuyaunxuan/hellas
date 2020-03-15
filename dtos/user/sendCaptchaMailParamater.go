package user


type SendCaptchaMailParameter struct {
	MailAddress string `json:"mail" binding:"required,email,max=50"`
}

