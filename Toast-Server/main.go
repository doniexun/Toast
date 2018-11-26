package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/doniexun/Toast/Toast-Server/routers"
	_ "github.com/doniexun/Toast/Toast-Server/models"
	"github.com/astaxie/beego"
)

//初始化
func init(){
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
    	dbuser := beego.AppConfig.String("dbuser")
    	dbpassword := beego.AppConfig.String("dbpassword")
    	dbname := beego.AppConfig.String("dbname")

    	//注册mysql Driver
    	orm.RegisterDriver("mysql", orm.DRMySQL)
	//构造conn连接
	//用户名:密码@tcp(url地址)/数据库
    	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
    	//注册数据库连接
    	orm.RegisterDataBase("default", "mysql", conn)
	
    	fmt.Printf("数据库连接成功！%s\n", conn)  
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)  // true 改成false，如果表存在则会给出提示，如果改成false则不会提示 ， 这句话没有会报主键不存在的错误

    	//o := orm.NewOrm()
    	//o.Using("default") // 默认使用 default，你可以指定为其他数据库

    	//启用Session
    	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
