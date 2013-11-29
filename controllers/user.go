package controllers

import (
 "fmt"
 "io"
 "crypto/md5"
 . "dilu/models"
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
	tomd5 := md5.New()
    username := this.GetString("username")
    password := this.GetString("password")
	io.WriteString(tomd5,password)
	pwd := tomd5.Sum(nil)
    u := User{}
  //  err := o.Read(&u,"username")
	qs := o.QueryTable("user")
	qs.Filter("username",username)
	qs.Filter("password",password)
	num,err := qs.Filter("username",username).Filter("password",password).All(&u)
    if err != nil{
    fmt.Printf("error %s",err)
    }else if(num<=0){
    fmt.Printf("not exist")
    }else{
		fmt.Printf("exist %d %x",num,pwd)
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

