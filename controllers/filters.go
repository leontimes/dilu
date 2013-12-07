package controllers
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)
type FilterController struct{
    beego.Controller
}
/*var FilterUser = func(ctx *context.Context) {
	var globalSessions *session.Manager
	globalSessions, _ = NewManager("memory","gosessionid",3600)
	sess := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	acc := sess.Get("account")
	if acc == nil{
		ctx.Redirect(302, "/login")
	}
	if ctx.Input.CruSession == nil {
		ctx.Input.CruSession = beego.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	}
	_, ok := ctx.Input.Session("account").(int)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}        */
func (this *FilterController) FilterUsers(ctx *context.Context){
	// this.GetSession("account")
	//	if account == nil {
		//this.SetSession("asta", int(1))
		//this.Data["num"] = 0
      //  this.Redirect("/login",302)
	_, ok := ctx.Input.Session("account").(int)
	if !ok {
		ctx.Redirect(302, "/login")
	}
	 ctx.Redirect(302, "/login")
	//	}
}
