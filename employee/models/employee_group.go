package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type UtEmployeeGroup struct {
	Id          int
	Name        string
	Description string
	CreateTime  time.Time
	UtEmployee  []*UtEmployee `orm:"reverse(many)"`
}

// JudgeGroupById
func JudgeGroupById(gid int) *UtEmployeeGroup {
	o := orm.NewOrm()
	group := &UtEmployeeGroup{Id: gid}
	if err := o.Read(group); err == nil { return group }
	return nil
}

// JudgeGroupByName
func JudgeGroupExistByName(name string) bool {
	o := orm.NewOrm()
	if exist := o.QueryTable("UtEmployeeGroup").Filter("Name", name).Exist(); exist {
		return false
	}
	return true
}


// InsertNewGroup插入组
func InsertNewGroup(name string) (bool, *UtEmployeeGroup)  {
	group := new(UtEmployeeGroup)
	o := orm.NewOrm()
	if exist := o.QueryTable("UtEmployeeGroup").Filter("Name", name).Exist(); exist {
		return false, nil
	}

	group.CreateTime = time.Now()
	group.Name = name
	if _, err := o.Insert(group); err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, group
}


// DeleteGroup移除组
func DeleteGroup(group *UtEmployeeGroup) bool {
	o := orm.NewOrm()
	if _, err := o.Delete(group); err != nil {
		return false
	}
	return true
}

// UpdateGroup 更新组
func UpdateGroup(group *UtEmployeeGroup) bool {
	o := orm.NewOrm()
	if _, err := o.Update(group); err != nil {
		return false
	}
	return true
}

// GetGroupDetail 通过组id获得组详情
func GetGroupDetail(group *UtEmployeeGroup) *UtEmployeeGroup {
	o := orm.NewOrm()
	if err := o.QueryTable("UtEmployeeGroup").Filter("Id", group.Id).RelatedSel().One(group); err == nil {
		if _, err := o.LoadRelated(group, "UtEmployee"); err != nil {
			return nil
		}
		return group
	}
	return nil
}


// GroupListPageOut 分页
func GroupListPageOut(CurrPage, PageSize int) ([]*UtEmployeeGroup, int64, error) {
	currPage := 1
	pageSize := 10
	var (
		count 	int64			// 总条数
		err 	error			// 错误
		res		[]*UtEmployeeGroup	// 结果
	)
	if CurrPage != -1 { currPage = CurrPage }
	if PageSize != -1 { pageSize = PageSize }
	o := orm.NewOrm()
	// 计算偏移
	offset := (currPage - 1) * pageSize
	qs := o.QueryTable("UtEmployeeGroup")
	count, err = qs.Count()
	if err == nil {
		// 查询
		_, err = qs.OrderBy("-CreateTime").Limit(pageSize, offset).RelatedSel().All(&res)
	}
	return res, count, err
}

//AddAgent2Group
func AddAgent2Group(group *UtEmployeeGroup, aid int) bool {
	o := orm.NewOrm()
	// 判断员工是否存在
	agent := &UtEmployee{Id: aid}
	if err := o.Read(agent); err != nil {
		return false
	}
	m2m := o.QueryM2M(agent, "UtEmployeeGroup")
	_, err := m2m.Add(group)
	if err != nil {
		return false
	}
	return true
}

// RemoveAgentFromGroup
func RemoveAgentFromGroup(gid, aid int) bool{
	o := orm.NewOrm()
	//rel := &UtRelEmployeeGroup{}

	if _, err := o.QueryTable("UtRelEmployeeGroup").Filter("UtEmployee", aid).
		Filter("UtEmployeeGroup", gid).RelatedSel().Delete(); err != nil {
			return false
	}

	return true
}