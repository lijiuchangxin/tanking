package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	. "time"
)

type UtEmployee struct {
	Id                 int
	Name               string
	Email              string
	Alias              string
	telephone          string
	duty               *UtDuty `orm:"rel(fk)"`
	ImAbility          int
	SipAccount         string
	SipPassword        string
	Avatar             string
	WorkId             string
	Department         *UtDepartment `orm:"rel(fk)"`
	AgentCalloutNumber string
	Status             string
	CreateTime         Time
	LoginTime          Time
	Profile            *UtProfile `orm:"rel(fk)"`
	Availability       string
	ImWelcome          string
	Language           string
	UtRole             []*UtRole          `orm:"rel(m2m);rel_through(gowebsite/apps/admin/employee/models.UtRelEmployeeRole)"`
	UtEmployeeGroup    []*UtEmployeeGroup `orm:"rel(m2m);rel_through(gowebsite/apps/admin/employee/models.UtRelEmployeeGroup)"`

}

//type UtEmployeeGroup struct {
//	Id          int
//	Name        string
//	Description string
//	CreateTime  Time
//	UtEmployee  []*UtEmployee `orm:"reverse(many)"`
//}

type UtDuty struct {
	Id          int
	name        string
	Description string
}

//type UtDepartment struct {
//	Id          string `orm:"pk"`
//	Name        string
//	Description string
//}

type UtProfile struct {
	Id          int
	Name        string
	Description string
}

type UtRelEmployeeGroup struct {
	Id              int
	UtEmployee      *UtEmployee      `orm:"rel(fk)"`
	UtEmployeeGroup *UtEmployeeGroup `orm:"rel(fk)"`
}

type UtRelEmployeeRole struct {
	Id         int
	UtEmployee *UtEmployee `orm:"rel(fk)"`
	UtRole     *UtRole     `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(UtProfile), new(UtDepartment), new(UtDuty), new(UtEmployeeGroup), new(UtRelEmployeeGroup), new(UtRelEmployeeRole), new(UtEmployee))
}

//query  rel-all
func QueryEmployee() ([]*UtEmployee, error) {
	o := orm.NewOrm()
	var utEmployees []*UtEmployee
	_, err := o.QueryTable("UtEmployee").All(&utEmployees)
	for i, utEmployee := range utEmployees {

		_, err := o.LoadRelated(utEmployee, "UtRole")
		if err != nil {
			logs.Error(err)
		}
		utEmployees[i] = utEmployee
	}

	for i, utEmployee := range utEmployees {
		o.LoadRelated(utEmployee, "UtEmployeeGroup")
		if err != nil {
			logs.Error(err)
		}
		utEmployees[i] = utEmployee
	}
	fmt.Println(utEmployees[0])
	fmt.Println(utEmployees[1])
	return utEmployees, err

}

//add empoy  rel2 (group 、role)
func AddEmployee(utEmpoyee UtEmployee) error {
	o := orm.NewOrm()
	_, err := o.QueryM2M(&utEmpoyee, "utRole").Add()
	if err != nil {
		logs.Error(err)
	}

	return err
}
