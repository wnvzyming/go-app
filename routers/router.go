package routers

import (
	"gocms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	//beego.AutoRouter(&controllers.AdminController{})
	beego.Router("/home", &controllers.HomeController{}, "*:Index")

	beego.AutoRouter(&controllers.AdminController{}) // 把需要的路由注册到自动路由中

	//	beego.Router("/admin/getList", &controllers.AdminController{})
	//	beego.Router("/admin/index", &controllers.AdminController{})

	//	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
}
