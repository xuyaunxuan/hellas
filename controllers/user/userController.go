package user

import (
	"hellas/controllers"
)

type Controller struct {
	controllers.BaseController
}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

func (c *Controller) CreateUser() {
	c.ParseToken()
	var result ShortResult

	result.UrlLong = "languid"
	c.Data["json"] = result
	c.ServeJSON()
}