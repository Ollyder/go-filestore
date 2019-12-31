package controllers
import (
	"github.com/astaxie/beego"
	"go-filestore/data"
	"fmt"
	// user "go-filestore/model/loginandregister"
	
)

type LoginController struct {
	beego.Controller
}

// Post : 登录操作
func (this *LoginController) Post() {
	/*user := user.User{}
	if err := this.ParseForm(&user); err != nil {
		resp := resp.Fail(err.Error())
		this.Data["json"] = &resp
		this.ServeJSON()
		return
	}*/
	resp := data.Success(1)
	this.Data["json"] = &resp
	this.ServeJSON()
	
	// 比对账号与密码

	// 登录成功写入session

	// 返回成功信息

}

func (this *LoginController) Get() {
	fmt.Println("!!!")
}