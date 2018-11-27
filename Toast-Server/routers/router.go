package routers

import (
	"github.com/doniexun/Toast/Toast-Server/controllers/web"
	"github.com/astaxie/beego"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/profile", &controllers.LoginController{}, "get:Profile")
}
