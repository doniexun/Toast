package controllers

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/doniexun/Toast/Toast-Server/models"
	"github.com/doniexun/Toast/Toast-Server/utils"
)

type LoginController struct {
	beego.Controller
}

//登录页面
func (c *LoginController) Get() {
	c.TplName = "web/login.html"
}

//登录功能
func (c *LoginController) Post() {
	name := c.GetString("name")
	pwd := c.GetString("pwd")

	var user models.User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("name", name).One(&user)
	
	isLogin := false
	fmt.Println(user.Password)
	fmt.Println(pwd)
	fmt.Println(utils.Md5([]byte(pwd)))
	if user.Id != 0 && user.Password == utils.Md5([]byte(pwd)) {
		isLogin = true

		c.SetSession("loginuser", int64(user.Id))
		fmt.Println("当前的session:")
		fmt.Println(c.CruSession)

		user.LastLoginIP = c.Ctx.Request.RemoteAddr
		user.LastLoginTime = time.Now()

	}else {	// 用户名或密码错误
		isLogin = false
	}

	c.Data["json"]=map[string]interface{}{"isLogin" : isLogin}
	c.ServeJSON()
}


func (c *LoginController) Profile() {
	id := c.GetSession("loginuser")
	if id_int64, ok := id.(int64); ok && id_int64 > 0{
		var user = models.User { Id : id.(int64) }
		o := orm.NewOrm()
		o.Read(&user)
		c.Data["user"] = user
		c.TplName = "web/profile.html"
	} else {
		c.Data["errMsg"] = "你还没登录，请先登录！"  // TODO 3秒后转到登录页面
		c.TplName = "web/error.html"
	}
}
//退出
type LogoutController struct {
	beego.Controller
}

//登录退出功能
func (c *LogoutController) Post() {
	v := c.GetSession("loginuser")
	isLogout := false
	if v != nil {
	  //删除指定的session	
	  c.DelSession("loginuser")

	  //销毁全部的session
	  c.DestroySession()
	  isLogout = true
	  
	 fmt.Println("当前的session:")
	 fmt.Println(c.CruSession)
	}

	c.Data["json"]=map[string]interface{}{"isLogout" : isLogout}
	c.ServeJSON()

	c.StopRun()
}
