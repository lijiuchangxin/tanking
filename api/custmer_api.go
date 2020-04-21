package api

import (
	"customer_managenment/models"
)

// 校验输入参数是否正确
type Verify interface {
	VerifyInputPara() bool
}

// 公共返回体
type CommonResponse struct {
	// code： 通常 0 表示正确， 1 请求错误，  2 内部错误， 其他值自定义
	Code  int          `json:"code"`
	Msg   string       `json:"msg"`
}

// 新增客户返回体
type ResponseNewCustomer struct {
	CommonResponse
	//Id 	  int		   	`json:"id"`
	UtCustomer  struct{
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
		Alters 				interface{}	`json:"alters"`
		FollowUp 			interface{}	`json:"follow_up"`
	}
}

// 新增客户请求体
type RequestNewCustomer struct {
	models.UtCustomer
}

// 删除客户请求体
type RequestDelCustomer struct {
	Id				 int		`json:"id"`
	OpenApiToken	string		`json:"open_api_token"`

}

// 删除客户返回体
type ResponseDelCustomer struct {
	CommonResponse
	Data struct{
		Id				int			`json:"id"`
		OpenApiToken	string		`json:"open_api_token"`
	}								`json:"data"`
}

// 新增客户跟进请求体
type RequestNewFollow struct {
	CustomerId	int				`json:"customer_id"`
	models.CustomerFollowUp
}

// 新增客户跟进返回体
type ResponseNewFollow struct {
	CommonResponse
	CustomerFollowUp 	`json:"customer_follow_up"`
}

// 删除跟进请求体
type RequestDelFollow struct {
	Id				 int		`json:"id"`

}

// 删除跟进返回体
type ResponseDelFollow struct {
	CommonResponse
	Data struct{
		Id				int			`json:"id"`
		CustomerId		int			`json:"customer_id"`
	}								`json:"data"`
}

// 客户详情请求体
type RequestShowCustomer struct {
	CustomerId	int  `json:"customer_id"`
}

// 客户详情返回体
type ResponseShowCustomer struct {
	ResponseNewCustomer
}

// 修改客户请求体
type RequestUpdateCustomer struct {
	CustomerId	int  		`json:"customer_id"`
	Tag			interface{}  	`json:"tag"`
	IsVip		interface{}		`json:"is_vip"`
	Desc		interface{}		`json:"desc"`
	TelPhone	interface{}		`json:"tel_phone"`
	CellPhone	interface{}		`json:"cell_phone"`
}

// 修改客户响应体
type ResponseUpdateCustomer struct {
	CommonResponse

}

//  跟进详情
type CustomerFollowUp struct {
	Id 				int			`json:"id"`
	Content			string		`json:"content"`
	CreateAt		int			`json:"create_at"`
	UserId			int			`json:"user_id"`
	UserAvatar		string		`json:"user_avatar"`
	UserNickName	string		`json:"user_nick_name"`
}

// 变更详情
type CustomerAlter struct {
	Id 				int			`json:"id"`
	UserId			int			`json:"user_id"`
	UserNickName 	string		`json:"user_nike_name"`
	AlterTime		int			`json:"alter_time"`
	Summary			string		`json:"summary"`
}

// 校验新增客户跟进入参合法性
func (request *RequestNewFollow)VerifyInputPara() bool {
	// customer必须>=0, userId>=0, content不能为空
	if request.CustomerId <= 0 || request.UserId <= 0 || request.Content == "" { return false }
	return true
}

// 校验新增客户请求体入参合法性
func (request *RequestNewCustomer)VerifyInputPara() bool {
	// 手动添加客户name为必填，客户端添加客户OpenApiToken为必填
	if request.OpenApiToken == "" && request.CustomerNikeName == "" { return false }
	return true
}

// 校验删除客户请求体入参合法性
func (request *RequestDelCustomer)VerifyInputPara() bool {
	// 所删除额客户的id必须为正整数或apiToken部位空
	if request.Id <= 0 && request.OpenApiToken == "" { return false }
	return true
}

// 校验删除跟进入参数合法性校验
func (request *RequestDelFollow)VerifyInputPara() bool {
	//  删除的客户跟进id必须>0
	if request.Id  <= 0 { return false }
	return true
}

// 查询客户详情入参数合法性校验
func (request *RequestShowCustomer)VerifyInputPara() bool {
	//  删除的客户跟进id必须>0
	if request.CustomerId  <= 0 { return false }
	return true
}

// 修改客户入参数合法性校验
func (request *RequestUpdateCustomer)VerifyInputPara() bool {
	//  删除的客户id必须>0
	if request.CustomerId  <= 0 { return false }
	return true
}

