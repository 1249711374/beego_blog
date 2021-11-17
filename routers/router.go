package routers

import (
	"beego_blog/controllers"
	"beego_blog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	// /index2.html
	beego.Router("/index:page:int.html", &controllers.MainController{}, "*:Index")

	beego.Router("/article/:id:int", &controllers.MainController{}, "*:Show")
	// 关于我
	beego.Router("/about.html", &controllers.MainController{}, "*:About")
	// 成长录
	beego.Router("/life.html", &controllers.MainController{}, "*:BlogList")

	//-----------------------------------------------------账户管理-----------------------------------------------
	//http://localhost:8080/admin/login
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")

}
