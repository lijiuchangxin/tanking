package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UtRole struct {
	Id          int
	Name        string
	Description string
	//UtOperation []*UtOperation `orm:"rel(m2m);rel_table(ut_re1l_role_operation)"`
	UtOperation []*UtOperation `orm:"rel(m2m);rel_through(gowebsite/apps/admin/employee/models.UtRelRoleOperation)"`
	UtEmployee  []*UtEmployee  `orm:"reverse(many)"`
}

type UtOperation struct {
	Id     int
	Name   string
	Belong int
	UtRole []*UtRole `orm:"reverse(many)"`

}

type UtRelRoleOperation struct {
	Id          int
	UtRole      *UtRole      `orm:"rel(fk)"`
	UtOperation *UtOperation `orm:"rel(fk)"`
}

//func (relRoleOperation *UtRelRoleOperation) TableName() string {
//	return "ut_rel_role_operation"
//}

func (UtOperation *UtOperation) TableName() string {
	return "ut_operation"
}

func init() {
	orm.RegisterModel(new(UtRole), new(UtOperation), new(UtRelRoleOperation))
	orm.Debug = true

}

// add ut_role	ut_operation	ut_rel_role_operation
//完成 rel operation
func AddRole(utRole UtRole) (int, error) {
	o := orm.NewOrm()
	fmt.Println("utRole", &utRole)
	//utRole1:=UtRole{Id:0,Name:"test5",Description:"desc,,,"}
	//fmt.Println("utRole1",utRole1)
	num, errRol := o.Insert(&utRole)

	//p,errRel:=o.Raw("insert into ut_rel_role_operation (role_id ,operation_id ) values(?,?)").Prepare()
	//_,err:=p.Exec(3,2)
	//_,err=p.Exec(3,4)
	//p.Close()
	//utOperation:=utRole.UtOperation
	if errRol != nil {
		logs.Error("add failed,", errRol)
	} else {
		qs := o.QueryM2M(&utRole, "UtOperation")
		for _, operation := range utRole.UtOperation {
			_, errRol = qs.Add(operation)
			if errRol != nil {
				logs.Error(errRol)
				break
			}
		}
	}

	return int(num), errRol
}

//query  ut_rel_role_operation
func QueryRole() ([]UtRole, error) {
	o := orm.NewOrm()
	//UtRoles :=[]UtRole{UtRole{Id:1},UtRole{Id:2}}
	var UtRoles []UtRole
	//UtRoles:=[]UtRole(UtRole{Id:1},UtRole{Id:2})
	//关联添加
	//num,err:=o.QueryM2M(&UtRole{Id:2},"utOperation").Add(&UtOperation{Id:1})
	//result,err:=o.Raw("select r.id Id ,r.name roleName,r.description,o.id operationId,o.name operationName,o.belong " +
	//	"from ut_rel_role_operation rel left join ut_role r   on r.id=rel.role_id left join ut_operation o on rel.operation_id =o.id").QueryRows(&UtRoles)
	num, err := o.QueryTable("UtRole").OrderBy("Id").RelatedSel().All(&UtRoles)
	//num, err = o.LoadRelated(&UtRoles,"UtOperation")
	for i, roleTmp := range UtRoles {
		num, err = o.LoadRelated(&roleTmp, "UtOperation")
		fmt.Println(num)
		UtRoles[i] = roleTmp
	}
	if err != nil {
		fmt.Println("query failed")
		logs.Error("query failed,", err)
	}
	fmt.Println(num)
	fmt.Println(UtRoles)

	return UtRoles, err
}

//按id查询role，关联operation
func QueryRoleById(id int) (UtRole, error) {
	o := orm.NewOrm()
	utRole := UtRole{Id: id}
	//var utRole UtRole
	//o.Read(&utRole)
	o.QueryTable(new(UtRole)).Filter("Id", id).RelatedSel().One(&utRole)
	num, err := o.LoadRelated(&utRole, "UtOperation")
	if err != nil {
		fmt.Println("query failed")
		logs.Error("query failed,", err)
	}
	fmt.Println(num)
	fmt.Println(utRole)
	return utRole, err
}

//rel 先删后增
func Update(utRole UtRole) (int, error) {
	o := orm.NewOrm()
	//o.Read(&utRole,"Id")
	fmt.Println(utRole)
	num, err := o.Update(&utRole)
	if err != nil {
		logs.Error("update role failed,", err)
	} else {
		//utRoleOperation:=UtRelRoleOperation{UtRole}
		num, err = o.Delete(&UtRelRoleOperation{UtRole: &utRole}, "UtRole")
		for _, utOperation := range utRole.UtOperation {
			utRoleOperation := UtRelRoleOperation{UtRole: &utRole, UtOperation: utOperation}
			num, err = o.Insert(&utRoleOperation)
		}
	}
	return int(num), err
}

//注意不存的记录或已经删除的记录，需要返回提示
//完成 delete rel_role_operation
func Delete(utRole UtRole) (int, error) {
	o := orm.NewOrm()
	//num, err :=o.QueryM2M(&utRole,"utOperation").Remove(UtOperation{Id:utRole.Id})
	utRelRoleOperation := UtRelRoleOperation{UtRole: &utRole}
	num, err := o.Delete(&utRelRoleOperation, "utRole")
	if err != nil {
		logs.Error("relevant delete failed,", err)
	} else {
		num, err = o.Delete(&utRole)
		if err != nil {
			logs.Error("delete failed,", err)
		}
	}

	return int(num), err
}
