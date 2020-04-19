package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"reflect"
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
}


type CustomerFollowUp struct {
	Id 				int			`json:"id"`
	CreateAt		int			`json:"create_at"`
	UpdatedAt		int			`json:"updated_at"`
	CustomerId		int			`json:"customer_id"`
	FeedType		string		`json:"feed_type"`
	UserId			int			`json:"user_id"`
	UserAvatar		string		`json:"user_avatar"`
	UserNickName	string		`json:"user_nick_name"`
}


type CustomerAlteration struct {
	Id 				int			`json:"id"`
	CustomerId		int			`json:"customer_id"`
	UserId			int			`json:"user_id"`
	UserNickName 	string		`json:"user_nike_name"`
	AlterTime		int			`json:"alter_time"`
	Summary			string		`json:"summary"`
	FeedType		string		`json:"feed_type"`
}


func CreateCustomer(customer *UtCustomer) error {
	o := orm.NewOrm()
	if _, err := o.Insert(customer); err != nil { return errors.New("create customer error") }
	return nil
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


func DeleteCustomer(cid int) error {
	o := orm.NewOrm()
	if _, err := o.Delete(&UtCustomer{Id: cid}); err != nil { return errors.New("delete error") }
	return nil
}


func ShowCustomerDetail(cid int) (res map[string]interface{}) {
	o := orm.NewOrm()
	customer := &UtCustomer{Id:cid}
	if err := o.Read(customer); err != nil { return nil }
	return nil
}

func JudgeApiIsExists(api string) bool {
	o := orm.NewOrm()
	if num, _ := o.QueryTable("UtCustomer").Filter("OpenApiToken", api).Count(); num > 0 { return false}
	return true
}

