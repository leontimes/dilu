package controllers
import (
	"html/template"
	"net/http"
	"github.com/astaxie/beego"
)
func Page_not_found(rw http.ResponseWriter, r *http.Request){
	//t,_:= template.New("beegoerrortemp").ParseFiles(beego.ViewsPath+"/error/404.html")
	t,_:= template.ParseFiles(beego.ViewsPath+"/error/404.html")
	data :=make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
