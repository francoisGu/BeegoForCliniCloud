package controllers

import (
	"beegoProject/APIConnect"
	"github.com/astaxie/beego"
)

const (
	clientId     = "X-CHALLENGE-APP"
	clientSecret = "Y2xpbmljbG91ZGNoYWxsZW5nZXNlY3JldGtleQ=="
	email        = "test@clinicloud.com"
	password     = "pass123"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.TplName = "index.tpl"

	succcessLogin, _, loginContent := APIConnect.LoginToAPI(email, password, clientId, clientSecret)
	if succcessLogin {
		_, _, userContent := APIConnect.GetUser(loginContent["uuid"].(string),
			loginContent["access_token"].(string))

		_, _, sessionContent := APIConnect.GetSession(loginContent)
		c.Data["login"] = loginContent
		c.Data["user"] = userContent
		c.Data["session"] = sessionContent
		c.TplName = "index.tpl"
	}

}

func (c *UserController) Post() {

}
