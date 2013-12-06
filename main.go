/**
 * Created by Administrator on 13-11-28.
 */
package main

import (
	"fmt"
	"dilu/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
)

func init() {
	// register model
	//  orm.RegisterModel(new(User))

	// set default database
	// orm.RegisterDataBase("default", "mysql", "root:123456@/gomysql?charset=utf8", 30)
}
type Manager struct {
	cookieName  string     //private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}
var FilterUser = func(ctx *context.Context) {
	var globalSessions *session.Manager
	globalSessions, _ = NewManager("memory","gosessionid",3600)
	sess := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	acc := sess.Get("account")
	if acc == nil{
		ctx.Redirect(302, "/login")
	}
	/*_, ok := ctx.Input.Session("account").(int)
	if !ok {
		ctx.Redirect(302, "/login")
	}        */
}
func main() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/", &controllers.UserController{})
	beego.Router("/login", &controllers.UserController{},"*:Login")
	beego.AddFilter("/user/*","AfterStatic",FilterUser)
	beego.Run()
	fmt.Printf("visit:localhost:%s",beego.AppConfig.String("httpport"))
}
