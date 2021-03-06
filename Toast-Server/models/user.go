package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id		int64		`orm:"auto;pk"`
	Name		string		`orm:"unique;size(100)"`
	Nickname	string		`orm:"unique;size(100)"`
	Password	string		`orm:"size(100)"`
	Email		string		`orm:"unique;size(100)"`
	Gender		string		`orm:"size(2)"`
	RoleId		string		`orm:"size(100)"`
	Status		string		`orm:"size(3)"`
	Phone		string		`orm:"size(16)"`
	QQ		string		`orm:"size(20)"`
	Wechat		string		`orm:"size(100)"`
	Weibo		string		`orm:"size(100)"`
	RegisterTime	time.Time	`orm:"auto_now_add;type(datetime)"`
	RegisterIP	string		`orm:"size(130)"`
	LastLoginTime	time.Time	`orm:"auto_now_add;type(datetime)"`
	LastLoginIP	string		`orm:"size(130)"`
}


func init() {
	orm.RegisterModel(new(User))
}

