package controllers

import (
	. "customer_managenment/api"
	"customer_managenment/models"
	. "customer_managenment/tool"
	"encoding/json"
	"github.com/astaxie/beego"
)

// customer控制器
type CustomerController struct {
	beego.Controller
}


// 解析参数和校验参数,post
func (c *CustomerController) AnalysisAndVerify(request Verify) bool {
	if c.Ctx.Request.Method != "POST" {
		return false
	}

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
	request := new(RequestNewCustomer)
	response := new(ResponseNewCustomer)
	InitResponse(&response.CommonResponse)
	//response.Code = 1
	//response.Msg = "success"

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
				response.Code = 2
			} else {
				response.Code = 0
				MapCustomerDetail(&response.CustomerDetail, &request.UtCustomer)
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
	request := new(RequestDelCustomer)
	response := new(ResponseDelCustomer)
	InitResponse(&response.CommonResponse)
	//response.Code = 1
	//response.Msg = "success"
	// 解析校验输入的参数
	if res := c.AnalysisAndVerify(request); res {
		// 通过apiToken删除客户，通过id删除客户
		// 如果id存在，则通过id删除客户
		// 如果apiToken存在，则通过token删除客户
		if  !models.JudgeIsExists("UtCustomer", "Id", request.Id) || (request.OpenApiToken != "" &&
			!models.JudgeIsExists("UtCustomer", "OpenApiToken", request.OpenApiToken)) {
			Logs.Info("customer", request.Id, request.OpenApiToken, "not exist")
			response.Msg = "deleted customer does not exist"
		} else {
			if !models.RemoveCustomer(request.Id, request.OpenApiToken) {
				Logs.Error("delete customer", request.Id, request.OpenApiToken, "fail")
				response.Msg = "failed to delete customer"
				response.Code = 2
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
	InitResponse(&response.CommonResponse)
	//response.Code = 1
	//response.Msg = "success"
	// 解析校验输入的参数
	if res := c.AnalysisAndVerify(request); res {
		// TODO 判断的权限以及合法性
		//判断是否存在customer
		if customer := models.GetCustomerById(request.CustomerId); customer == nil {
			Logs.Info("ut_customer_id not exist")
			response.Msg = "the customer to be followed up dose not exist"
		} else {
			request.Customer = customer
			if err := models.InsertCustomerFollow(&request.CustomerFollowUp); err != nil {
				// 新数据插入数据库失败
				Logs.Error("new customer follow up failed, because", err)
				response.Msg = "new customer follow up failed"
				response.Code = 2
			} else {
				// 添加成功，更新返回题
				Logs.Info("new follow up success")
				response.Code = 0
				MapCustomerFollow(&response.CustomerFollowUp, &request.CustomerFollowUp)
			}
		}
	} else {
		response.Msg = "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}

// DeleteCustomerFollow 删除客户跟进
func (c *CustomerController)DeleteCustomerFollow() {
	request := new(RequestDelFollow)
	response := new(ResponseDelFollow)
	InitResponse(&response.CommonResponse)
	//response.Code = 1
	//response.Msg = "success"
	// 解析校验输入的参数
	if res := c.AnalysisAndVerify(request); res {
		// 通过跟进的id判断跟进是否存在
		if !models.JudgeIsExists("CustomerFollowUp", "Id", request.Id) {
			Logs.Info("customer follow up", request.Id, "not exist")
			response.Msg = "deleted customer follow up does not exist"
		} else {
			custerID, res := models.RemoveCustomerFollow(request.Id)
			if !res {
				Logs.Error("delete customer follow up", request.Id, "fail")
				response.Msg = "failed to delete customer follow up"
				response.Code = 2
			} else {
				Logs.Info("delete", request.Id, "success")
				response.Code = 0
				response.Data.Id = request.Id
				response.Data.CustomerId = custerID
			}
		}
	} else {
		response.Msg = "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}

// ShowCustomerDetail 展示客户详情，包括跟进，变更
func (c *CustomerController)ShowCustomerDetail() {
	request := new(RequestShowCustomer)
	response := new(ResponseShowCustomer)
	InitResponse(&response.CommonResponse)
	//response.Code = 1
	//response.Msg = "success"
	if res, err := c.GetInt("customer_id"); err == nil {
		request.CustomerId = res
		if res := request.VerifyInputPara(); res {
			if res := models.GetCustomerById(request.CustomerId); res == nil {
				Logs.Info("customer dose not exist, customer_id:", request.CustomerId)
				response.Msg = "customer dose not exist"
			} else {
				response.Code = 0
				MapCustomerDetail(&response.CustomerDetail, res)
				//response.UtCustomer = *res
			}
		} else {
			response.Msg = "incoming parameter error"
		}
	} else {
		response.Msg = "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}

// UpdateCustomer 修改客户详情
func (c *CustomerController)UpdateCustomer() {
	request := new(RequestUpdateCustomer)
	response := new(ResponseUpdateCustomer)
	InitResponse(&response.CommonResponse)
	//response.Code = 1
	//response.Msg = "success"
	// 校验参数是否正确
	if res := c.AnalysisAndVerify(request); res {
		// 判断customer是否存在
		if !models.JudgeIsExists("UtCustomer", "Id", request.CustomerId) {
			Logs.Info("ut_customer_id not exist")
			response.Msg = "update customer failed, because customer_id dose not exist"
		} else {
			// 获取入参映射
			res := GetUpdateCustomerMap(request)
			if err := models.UpdateCustomer(request.CustomerId, res); err == nil {
				response.Code = 0
			} else {
				Logs.Info("update customer failed")
				response.Msg = err.Error()
				response.Code = 2
			}
		}
	} else {
		response.Msg = "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}

//CustomerListByPage 分页查询列表详情
func (c *CustomerController)CustomerListByPage() {
	request := new(RequestCustomerList)
	response := new(ResponseCustomerList)
	InitResponse(&response.CommonResponse)

	request.CurrPage, _ = c.GetInt("curr_page")
	request.PageSize, _ = c.GetInt("page_size")

	// 校验入参是否正确
	//if res := c.AnalysisAndVerify(request); res {
	if res := request.VerifyInputPara(); res {
		if res, count, err := models.CustomerListPageOut(request.CurrPage, request.PageSize); err != nil {
			response.Code = 2
			if count == 0 {
				response.Msg = "error in getting total number of customer details"
			} else {
				response.Msg = "failed to get customer details on page"
			}
		} else {
			m := make(map[int]*CustomerDetail)
			response.Customers = m
			for num, customer := range res {
				customerDetail := new(CustomerDetail)
				MapCustomerDetail(customerDetail, customer)
				response.Customers[num] = customerDetail
				response.Code = 0
				//response.CustomerList = append(response.CustomerList, customerDetail)
			}
		}
	} else {
		response.Msg = "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}

// CustomerSearch
func (c *CustomerController)CustomerSearch()  {
	request := new(RequestSearchCustomer)
	response := new(ResponseSearchCustomer)
	InitResponse(&response.CommonResponse)

	// 校验参数是否正确
	if res := c.AnalysisAndVerify(request); res {
		if res, err := models.CustomerSearchByFiled(request.FiledName, request.Value); err != nil {
			response.Code = 2
			response.Msg = "customer query failed"
		} else {
			m := make(map[int]*CustomerDetail)
			response.Customers = m
			response.Code = 0
			for num, customer := range res {
				customerDetail := new(CustomerDetail)
				MapCustomerDetail(customerDetail, customer)
				response.Customers[num] = customerDetail
			}
		}
	} else {
		response.Msg = "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}