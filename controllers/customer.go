package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	. "customer_managenment/api"
	"customer_managenment/models"
	. "customer_managenment/tool"
)

type CustomerController struct {
	beego.Controller
}


// 解析参数和校验参数
func (c *CustomerController) AnalysisAndVerify(request Verify) bool {
	// 参数解析失败
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		Logs.Info("error parsing parameters")
		return false
	}
	// 校验参数失败
	if !request.VerifyInputPara() {
		Logs.Info("input para error")
		return false
	}
	return true
}


// NewCustomer 创建客户
func (c *CustomerController)CreateCustomer() {
	// 实例新增返回体，并默认初始值
	response := new(ResponseNewCustomer)
	response.Code = 1
	response.Msg = "success"
	request := new(RequestNewCustomer)
	// 解析校验输入的参数
	if res := c.AnalysisAndVerify(request); res {
		// 判断该 OpenApiToken 是否已有注册
		if  models.JudgeIsExists("UtCustomer", "OpenApiToken", request.OpenApiToken) &&
			request.OpenApiToken != "" {
			Logs.Info(request.OpenApiToken, "already exist, don't register again")
			response.Msg = "the customer is registered"
		} else {
			// 数据库新增客户
			if err := models.InsertCustomer(&request.UtCustomer); err != nil {
				Logs.Error(request.OpenApiToken, "register fail, because", err)
				response.Msg = "failed to register customer"
			} else {
				Logs.Info(request.OpenApiToken, "register success")
				response.Code = 0
				response.UtCustomer = request.UtCustomer
			}
		}
	} else {
		response.Msg = "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}


// DeleteCustomer 删除客户
func (c *CustomerController)DeleteCustomer() {
	response := new(ResponseDelCustomer)
	response.Code = 1
	response.Msg = "success"
	request := new(RequestDelCustomer)
	// 解析校验输入的参数
	if res := c.AnalysisAndVerify(request); res {
		// 通过apiToken删除客户，通过id删除客户
		// 如果id存在，则通过id删除客户
		// 如果apiToken存在，则通过token删除客户
		if  !models.JudgeIsExists("UtCustomer", "Id", request.Id) && (request.OpenApiToken != "" &&
			!models.JudgeIsExists("UtCustomer", "OpenApiToken", request.OpenApiToken)) {
			Logs.Info("customer", request.Id, request.OpenApiToken, "not exist")
			response.Msg = "deleted customer does not exist"
		} else {
			if !models.RemoveCustomer(request.Id, request.OpenApiToken) {
				Logs.Error("delete customer", request.Id, request.OpenApiToken, "fail")
				response.Msg = "failed to delete customer"
			} else {
				Logs.Info("delete", request.Id, request.OpenApiToken, "success")
				response.Code = 0
			}
		}
	} else {
		response.Msg = "incoming parameter error"
	}
	response.Data.OpenApiToken = request.OpenApiToken
	response.Data.Id = request.Id
	c.Data["json"] = response
	c.ServeJSON()
	return
}

// CreateCustomerFollow新增客户跟进
func (c *CustomerController)CreateCustomerFollow() {
	request := new(RequestNewFollow)
	response := new(ResponseNewFollow)
	response.Code = 1
	response.Msg = "success"
	// 解析校验输入的参数
	if res := c.AnalysisAndVerify(request); res {
		// TODO 判断的权限以及合法性
		if res := models.JudgeIsExists("UtCustomer", "Id", request.Customer.Id); !res {
			Logs.Info("ut_customer_id not exist")
			response.Msg = "the customer to be followed up dose not exist"
		} else {
			if err := models.InsertCustomerFollow(&request.CustomerFollowUp); err != nil {
				Logs.Error("new customer follow up failed, because", err)
				response.Msg = "new customer follow up failed"
			} else {
				Logs.Info("new follow up success")
				response.Code = 0
				response.CustomerFollowUp = request.CustomerFollowUp
			}
		}
	} else {
		response.Msg = "incoming parameter error"
	}
}

// DeleteCustomerFollow删除客户跟进
func (c *CustomerController)DeleteCustomerFollow() {

}
