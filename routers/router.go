package routers

import (
	"github.com/astaxie/beego"
	"hellas/controllers/user"
)

func init() {
	beego.Router("/create/user", &user.Controller{}, "get:CreateUser")
	beego.Router("/login", &user.Controller{}, "get:CreateUser")
}
