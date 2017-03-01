package routers

import (
	"beego-ex/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/join", &controllers.MainController{}, "get:Join")
	beego.Router("/user", &controllers.UserController{})
}
