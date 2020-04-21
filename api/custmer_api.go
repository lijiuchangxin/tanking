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
	models.UtCustomer	`json:"customer"`
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
	CustomerFollowUp struct{
		Id 				int			`json:"id"`
		Content			string		`json:"content"`
		CreateAt		int			`json:"create_at"`
		CustomerId		int			`json:"customer_id"`
		UserId			int			`json:"user_id"`
		UserAvatar		string		`json:"user_avatar"`
		UserNickName	string		`json:"user_nick_name"`
	}								`json:"customer_follow_up"`


	//models.CustomerFollowUp		`json:"customer_follow_up"`
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
		CustomerId		int		`json:"customer_id"`
	}								`json:"data"`
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

// 校验删除跟进入参数合法行
func (request *RequestDelFollow)VerifyInputPara() bool {
	//  删除的客户跟进id必须>0
	if request.Id  <= 0 { return false }
	return true
}



