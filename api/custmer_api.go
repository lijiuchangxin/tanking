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
	Id 	  int		   `json:"id"`
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

// 删除客户返回题
type ResponseDelCustomer struct {
	CommonResponse
	Data struct{
		Id				int			`json:"id"`
		OpenApiToken	string		`json:"open_api_token"`
	}								`json:"data"`
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



