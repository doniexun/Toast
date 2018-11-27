package controllers

import (
	"fmt"
	"strings"
	"github.com/astaxie/beego"
	"github.com/doniexun/Toast/Toast-Server/models"
	"github.com/doniexun/Toast/Toast-Server/utils"
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
	// 第一重处理：去掉收尾空格
	name := strings.TrimSpace(c.GetString("name"))
	nickname := strings.TrimSpace(c.GetString("nickname"))
	password := strings.TrimSpace(c.GetString("pwd"))
	email := strings.TrimSpace(c.GetString("email"))
	phone := strings.TrimSpace(c.GetString("phone"))
	qq := strings.TrimSpace(c.GetString("qq"))
	wechat := strings.TrimSpace(c.GetString("wechat"))
	
	// TODO 先不校验，直接写入数据库表	
	user := models.User{}
	user.Name = name
	user.Nickname = nickname
	user.Password = utils.Md5([]byte(password))
	user.Email = email
	user.Phone = phone
	user.QQ = qq
	user.Wechat = wechat
	user.RegisterIP = c.Ctx.Request.RemoteAddr
	user.LastLoginIP = c.Ctx.Request.RemoteAddr

	isSuccess := false
	o := orm.NewOrm()
	id, err := o.Insert(&user)			// 将用户信息插入数据表中
	fmt.Println("New user id:%d\n", id)
	if err == nil {
		isSuccess = true
		c.SetSession("loginuser", int64(user.Id))
		fmt.Println("当前的session:")
		fmt.Println(c.CruSession)
	}

	c.Data["json"]=map[string]interface{}{"isSuccess":isSuccess};
	c.ServeJSON();					// 只要有输出，后续代码就不会再运行
}


