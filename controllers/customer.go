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
func (c *CustomerController)NewCustomer() {
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
			response.Msg = "该客户已注册"
		} else {
			// 数据库新增客户
			if res, err := models.CreateCustomer(&request.UtCustomer); !res {
				Logs.Error(request.OpenApiToken, "register fail, because", err)
				response.Msg = "注册客户失败"
			} else {
				Logs.Info(request.OpenApiToken, "register success")
				response.Code = 0
			}
		}
	} else {
		response.Msg = "传入参数错误"
	}
	response.Id = request.Id
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
		if  !models.JudgeIsExists("UtCustomer", "Id", request.Id) || (request.OpenApiToken != "" &&
			!models.JudgeIsExists("UtCustomer", "OpenApiToken", request.OpenApiToken)) {
			Logs.Info("customer", request.Id, request.OpenApiToken, "not exist")
			response.Msg = "删除的客户不存在"
		} else {
			if !models.RemoveCustomer(request.Id, request.OpenApiToken) {
				Logs.Error("delete customer", request.Id, request.OpenApiToken, "fail")
				response.Msg = "删除客户失败"
			} else {
				Logs.Info("delete", request.Id, request.OpenApiToken, "success")
				response.Code = 0
			}
		}
	} else {
		response.Msg = "传入参数错误"
	}
	response.Data.OpenApiToken = request.OpenApiToken
	response.Data.Id = request.Id
	c.Data["json"] = response
	c.ServeJSON()
	return
}

