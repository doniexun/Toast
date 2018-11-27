package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/doniexun/Toast/Toast-Server/models"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

//注册页面
func (c *RegisterController) Get() {
	c.TplName = "web/register.html"
}

//注册功能
func (c *RegisterController) Post() {
	name := c.GetString("name")
	nickname := c.GetString("nickname")
	pwd := c.GetString("pwd")
	email := c.GetString("email")
	phone := c.GetString("phone")
	qq := c.GetString("qq")
	wechat := c.GetString("wechat")
	
	// TODO 先不校验，直接写入数据库表	
	user := models.User{}
	user.Name = name
	user.Nickname = nickname
	user.Pwd = pwd
	user.Email = email
	user.Phone = phone
	user.QQ = qq
	user.Wechat = wechat
	user.RegisterIP = c.Ctx.Request.RemoteAddr
	user.LastLoginIP = c.Ctx.Request.RemoteAddr

	success := 0
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	fmt.Println("New user id:%d\n", id)
	if err == nil {
		success = 0
		c.SetSession("user", "adminuser")
		fmt.Println("当前的session:")
		fmt.Println(c.CruSession)
	} else {
		success = 1
	}

	c.Data["json"]=map[string]interface{}{"success":success};
	c.ServeJSON();
}

