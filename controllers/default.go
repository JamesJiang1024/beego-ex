package controllers

import (
	"github.com/astaxie/beego"
)

// MainController is a controller about user
type MainController struct {
	beego.Controller
}

// UserController is a controller about user
type UserController struct {
	beego.Controller
}

// Get function is to get user info
func (c *MainController) Get() {
	c.Data["Website"] = "wentian"
	c.Data["Email"] = "jwentian@redhat.com"
	c.TplName = "index.tpl"
}

// Get function is to get user info
func (c *UserController) Get() {
	c.Data["Website"] = "user"
	c.Data["Email"] = "user@redhat.com"
	c.TplName = "index.tpl"
}
