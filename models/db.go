package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//调用什么驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	//连接数据   ( 默认参数 ，mysql数据库 ，"数据库的用户名 ：数据库密码@tcp("+数据库地址+":"+数据库端口+")/库名？格式",默认参数）
	//orm.RegisterDataBase("default", "mysql", "root:124379685li!@tcp(127.0.0.1:3306)/customer?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "root:124379685@tcp(127.0.0.1:3306)/customer?charset=utf8")

	//注册model 建表
	orm.RegisterModel(new(UtCustomer), new(CustomerFollowUp), new(CustomerAlteration))

	orm.RunSyncdb("default", false, true)

}