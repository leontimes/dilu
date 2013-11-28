package controllers

import (
 "fmt"
 . "gweb/models"
 "github.com/astaxie/beego"
 "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)
type UserController struct{
   beego.Controller
}
func (this *UserController) Get(){
    this.Data["username"] = this.GetString("username")
   	this.Data["password"] = this.GetString("password")
	this.TplNames = "index.tpl"
}

func (this *UserController) Login(){
    o := orm.NewOrm()
    username := this.GetString("username")
    password := this.GetString("password")
    u := User{Username:username,Password:password}
    err := o.Read(&u,"username")
    if err == orm.ErrNoRows{
    fmt.Printf("not exist")
    }else{
    fmt.Printf("exist")
    }
    fmt.Printf(u.Username);
    this.TplNames = "index.tpl"
}

func (this *UserController) Post(){
    o := orm.NewOrm()
    username := this.GetString("username")
    password := this.GetString("password")
    user := User{Username:username,Password:password}
    // insert
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
    this.Data["username"] = this.GetString("username")
   	this.Data["password"] = this.GetString("password")
    this.TplNames = "index.tpl"


}

