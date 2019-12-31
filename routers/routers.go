package routers

import (
	"github.com/astaxie/beego"
	"go-filestore/controllers"
)

func init() {
	// 在此注册路由逻辑
	beego.Router("/login", &controllers.LoginController{})
}