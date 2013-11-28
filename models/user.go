package models

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)
type User struct {
    Id   int
    Username string `orm:"size(100)"`
    Password string
}
func init() {
// 将表定义注册到 orm 里
orm.RegisterModel(new(User))

// 读取配置参数
dbhost := beego.AppConfig.String("dbhost")
dbuser := beego.AppConfig.String("dbuser")
dbpwd := beego.AppConfig.String("dbpwd")
dbport := beego.AppConfig.String("dbport")
dbname := beego.AppConfig.String("dbname")

DSN := dbuser + ":" + dbpwd + "@(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

// 链接参数设置，最后的30，不知道是不是连接池数目
orm.RegisterDataBase("default", "mysql", DSN, 30)
orm.BootStrap() // 确保在所有
}