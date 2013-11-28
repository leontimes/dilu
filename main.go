/**
 * Created by Administrator on 13-11-28.
 */
package main

import (
	"gweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// register model
	//  orm.RegisterModel(new(User))

	// set default database
	// orm.RegisterDataBase("default", "mysql", "root:123456@/gomysql?charset=utf8", 30)
}

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/", &controllers.UserController{})
	beego.Run()
}
