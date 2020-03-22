package article

type DeleteParameter struct {
	ArticleId string `binding:"required,max=200"`
}
