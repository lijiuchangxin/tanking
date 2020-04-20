package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"time"
)

type UtCustomer struct {
	Id 					int			`json:"id"`
	CustomerNikeName 	string		`json:"customer_nike_name" form:"customer_nike_name "`
	Desc				string		`json:"desc"`
	Tag 				string		`json:"tag"`
	TelPhone 			string		`json:"tel_phone"`
	CellPhone 			string		`json:"cell_phone"`
	Email 				string		`json:"email"`
	IsVip 				int			`json:"is_vip"`
	Province  			string		`json:"province"`
	City 				string		`json:"city"`
	SourceChannel		string		`json:"source_channel"`
	CreateAt			int			`json:"create_at"`
	UpdatedAt			int			`json:"updated_at"`
	OrganizationName	string		`json:"organization _id"`
	OrganizationId		int			`json:"organization _name"`
	OwnerGroupId		int			`json:"owner_group_id"`
	OwnerGroupName		string		`json:"owner_group_name "`
	OwnerId				int			`json:"owner_id"`
	OwnerName			string		`json:"owner_name"`
	TicketCount			int			`json:"ticket_count"`
	LastContactAt		int			`json:"last_contact_at"`
	LastContactImAt		int			`json:"first_contact_at"`
	FirstContactAt		int			`json:"first_contact_im_at"`
	FirstContactImAt	int			`json:"last_contact_im_at"`
	OpenApiToken		string		`json:"open_api_token"`
	Alters 				[]*CustomerAlteration 	`orm:"reverse(many)"`
	FollowUps			[]*CustomerFollowUp		`orm:"reverse(many)"`
}


type CustomerFollowUp struct {
	Id 				int			`json:"id"`
	CreateAt		int			`json:"create_at"`
	UpdatedAt		int			`json:"updated_at"`
	//CustomerId		int			`json:"customer_id"`
	FeedType		string		`json:"feed_type"`
	UserId			int			`json:"user_id"`
	UserAvatar		string		`json:"user_avatar"`
	UserNickName	string		`json:"user_nick_name"`
	Customer		*UtCustomer	`orm:"rel(fk);cascade"`
}


type CustomerAlteration struct {
	Id 				int			`json:"id"`
	//CustomerId		int			`json:"customer_id"`
	UserId			int			`json:"user_id"`
	UserNickName 	string		`json:"user_nike_name"`
	AlterTime		int			`json:"alter_time"`
	Summary			string		`json:"summary"`
	FeedType		string		`json:"feed_type"`
	Customer		*UtCustomer	`orm:"rel(fk);cascade"`

}


// 判断是否存在
// table 查询的表名
// col 查询的列名
// res 查询的值
func JudgeIsExists(table, col string, res interface{}) bool {
	// false 不存在， true 存在
	o := orm.NewOrm()
	if exist := o.QueryTable(table).Filter(col, res).Exist(); exist { return true }
	return false
}

// CreateCustomer 新增客户详情
func CreateCustomer(customer *UtCustomer) (bool, error) {
	o := orm.NewOrm()
	createTime := int(time.Now().Unix())
	// 创建时间
	customer.CreateAt, customer.UpdatedAt = createTime, createTime
	// api用户昵称
	if customer.OpenApiToken != "" {
		customer.CustomerNikeName = fmt.Sprintf("API匿名客户(%s)", customer.OpenApiToken)
	}
	//_ = o.Begin()
	// 插入数据库
	if _, err := o.Insert(customer); err != nil {
		return false, err }
	alter := CustomerAlteration{
		//CustomerId:   customer.Id,
		Customer:	  customer,
		UserId:       customer.OwnerId,
		UserNickName: customer.OwnerName,
		AlterTime:    createTime,
		Summary:      "创建了客户",
	}
	// 创建配套跟进数据
	if err := CreateCustomerAlter(&alter); err != nil {
		//_ = o.Rollback()
		return false, err
	}
	//_ = o.Commit()
	return true, nil
}

// CreateCustomerAlter 新增客户跟进
func CreateCustomerAlter(alter *CustomerAlteration) error {
	o := orm.NewOrm()
	if _, err := o.Insert(alter); err != nil { return err }
	return nil
}

// RemoveCustomer 从客户表中删除客户
func RemoveCustomer(id int, token string) bool {
	o := orm.NewOrm()
	customer := UtCustomer{Id:id}
	// 通过id删除
	if err := o.Read(&customer); err == nil{
		if _, err := o.Delete(&customer); err != nil {
			return false }
	} else {
		// 通过token删除
		if _, err := o.QueryTable("UtCustomer").Filter("OpenApiToken", token).Delete(); err != nil {
			return false
		}
	}
	return true
}


func UpdateCustomer(cid int, paras map[string]interface{}) (err error) {
	o := orm.NewOrm()
	customer := &UtCustomer{Id:cid}
	if err = o.Read(customer); err != nil { return errors.New("read error") }

	if err = o.Begin(); err != nil {return errors.New("work error")}
	for key, value := range paras {
		customerValue := reflect.ValueOf(customer).Elem()
		if customerValue.FieldByName(key).Type() != reflect.TypeOf(value) {
			if err := o.Rollback(); err != nil { return  errors.New("rollback error") }
			return errors.New("para error")
		}
		customerValue.FieldByName(key).Set(reflect.ValueOf(value))
		if _, err = o.Update(customer, key); err != nil { return errors.New("update error") }
	}
	if err = o.Commit(); err != nil { return errors.New("commit error") }

	return
}


func DeleteCustomer(cid int) (bool, error) {
	o := orm.NewOrm()
	if _, err := o.Delete(&UtCustomer{Id: cid}); err != nil { return false, err }
	return true, nil
}


func ShowCustomerDetail(cid int) (res map[string]interface{}) {
	o := orm.NewOrm()
	customer := &UtCustomer{Id:cid}
	if err := o.Read(customer); err != nil { return nil }
	return nil
}


