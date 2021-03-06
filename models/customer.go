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
	CustomerNikeName 	string		`json:"customer_nike_name"`
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
	OrganizationName	string		`json:"organization_name "`
	OrganizationId		int			`json:"organization_id"`
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
	Alters 				[]*CustomerAlteration 	`orm:"reverse(many)" json:"alters"`
	FollowUps			[]*CustomerFollowUp		`orm:"reverse(many)" json:"follow_up"`
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
	Content			string		`json:"content"`
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


// 通过客户id得到dbCol
func GetCustomerById(id int) *UtCustomer {
	customer := &UtCustomer{Id:id}
	o := orm.NewOrm()
	if err := o.Read(customer); err != nil { return nil }
	// 传入customer为指针, 会直接修改
	GetFollowAndAlterByCid(customer, o)
	return customer
}

// 获得customer下的所有变更与跟进
func GetFollowAndAlterByCid	(customer *UtCustomer, o orm.Ormer)  {
	var alters []*CustomerAlteration
	var follows []*CustomerFollowUp
	// 获得所有的变更
	o.QueryTable("CustomerAlteration").Filter("Customer", customer.Id).RelatedSel().All(&alters)
	// 获得所有的跟进
	o.QueryTable("CustomerFollowUp").Filter("Customer", customer.Id).RelatedSel().All(&follows)
	customer.FollowUps = follows
	customer.Alters = alters
}

// InsertCustomer 新增客户详情
func InsertCustomer(customer *UtCustomer) error {
	o := orm.NewOrm()
	createTime := int(time.Now().Unix())
	// 创建时间
	customer.CreateAt, customer.UpdatedAt = createTime, createTime
	// api用户昵称
	if customer.OpenApiToken != "" && customer.CustomerNikeName == "" {
		customer.CustomerNikeName = fmt.Sprintf("API匿名客户(%s)", customer.OpenApiToken)
	}
	//_ = o.Begin()
	// 插入数据库
	if _, err := o.Insert(customer); err != nil {
		return err }
	alter := CustomerAlteration{
		//CustomerId:   customer.Id,
		Customer:	  customer,
		UserId:       customer.OwnerId,
		UserNickName: customer.OwnerName,
		AlterTime:    createTime,
		Summary:      "创建了客户",
	}
	// 创建配套跟进数据
	if err := InsertCustomerAlter(&alter); err != nil {
		//_ = o.Rollback()
		return err
	}
	customer.Alters = append(customer.Alters, &alter)
	//_ = o.Commit()
	return nil
}

// InsertCustomerAlter 新增客户变更
func InsertCustomerAlter(alter *CustomerAlteration) error {
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

// InsertCustomerFollow 插入客户跟进
func InsertCustomerFollow(follow *CustomerFollowUp) error {
	o := orm.NewOrm()
	createTime := int(time.Now().Unix())
	follow.CreateAt, follow.UpdatedAt = createTime, createTime
	if _, err := o.Insert(follow); err != nil { return err }
	return nil
}

// RemoveCustomerFollow 刪除客户跟进
func RemoveCustomerFollow(id int) (int, bool) {
	o := orm.NewOrm()
	follow := CustomerFollowUp{Id:id}
	// 通过id删除
	if err := o.Read(&follow); err == nil{
		if _, err := o.Delete(&follow); err == nil {
			return follow.Customer.Id, true
		}
	}
	return -1, false
}

// UpdateCustomer 更新客户详情
func UpdateCustomer(cid int, paras map[string]interface{}) error {
	o := orm.NewOrm()
	customer := &UtCustomer{Id:cid}
	if err := o.Read(customer); err != nil { return err }
	// 同过反射完成更新
	customerValue := reflect.ValueOf(customer).Elem()

	// 通过循环更新值与插入变更
	for key, value := range paras {
		if value != nil || value != "" {
			// 先存储旧值 后面要新增变更summary用的到
			oldValue := customerValue.FieldByName(key).String()
			// 如果新传入的值与旧值相同 跳出此次徐怒换
			if value == oldValue {
				continue
			}
			// 更新值
			if fmt.Sprintf("%T", value) == "float64" {
				if fmt.Sprintf("%v", value) == "1" {
					value = 1
				} else {
					value = 0
				}
			}
			customerValue.FieldByName(key).Set(reflect.ValueOf(value))
			if _, err := o.Update(customer, key); err != nil {
				return err
			} else {
				// 新增变更
				// 格式化summary
				if value == "" { value = "<空>" }
				if oldValue == "" { oldValue = "<空>"}
				follow := &CustomerAlteration{
					AlterTime:     int(time.Now().Unix()),
					// TODO 操作的客服，可能需要通过session获取
					//UserId:       0,
					//UserAvatar:   "",
					//UserNickName: "",
					Summary:      fmt.Sprintf("%s %s:%s ---> %s","lzs", key, oldValue, value),
					Customer:     customer,
				}
				if _, err := o.Insert(follow); err != nil { return errors.New("new customer follow up failed") }
			}
		}
	}
	// 更新时间
	customer.UpdatedAt = int(time.Now().Unix())
	if _, err := o.Update(customer, "UpdatedAt"); err != nil { return err}
	return nil
}

// CustomerListPageOut 分条从数据库中取数据
func CustomerListPageOut(CurrPage, PageSize int) ([]*UtCustomer, int64, error) {
	currPage := 1
	pageSize := 10
	var (
		count 	int64			// 总条数
		err 	error			// 错误
		res		[]*UtCustomer	// 结果
	)
	if CurrPage != -1 { currPage = CurrPage }
	if PageSize != -1 { pageSize = PageSize }
	o := orm.NewOrm()
	// 计算偏移
	offset := (currPage - 1) * pageSize
	qs := o.QueryTable("UtCustomer")
	count, err = qs.Count()
	if err == nil {
		// 查询
		_, err = qs.OrderBy("-CreateAt").Limit(pageSize, offset).RelatedSel().All(&res)
	}
	return res, count, err
}


// CustomerSearchByFiled 实现简单数据查找 下期再完善
func CustomerSearchByFiled(filed, value string) ([]*UtCustomer, error) {
	var res []*UtCustomer
	o := orm.NewOrm()
	qs := o.QueryTable("UtCustomer")
	_, err := qs.Filter(filed, value).All(&res)
	return res, err
}




