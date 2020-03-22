package article

type DeleteParameter struct {
	ArticlePath string `binding:"required,max=200"`
}
