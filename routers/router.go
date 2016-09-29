package routers

import (
	"beegoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{})
	// beego.Post("/", &controllers.MainController{})
}
