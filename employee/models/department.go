package models

import "github.com/astaxie/beego/orm"

type UtDepartment struct {
	Id          string `orm:"pk"`
	Name        string
	Description string
}


// JudgeGroupByName
func JudgeDepartExistByName(name string) bool {
	o := orm.NewOrm()
	if exist := o.QueryTable("UtEmployeeGroup").Filter("Name", name).Exist(); exist {
		return false
	}
	return true
}