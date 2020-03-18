package article

type SubscribeParameter struct {
	Title string `binding:"required,max=200"`
	Tag string `binding:"max=100"`
	IsPrivate bool `binding:""`
	Article string `binding:""`
	ArticleId string `binding:""`
}
