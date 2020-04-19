package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"lzsc/models"
)

type CustomerController struct {
	beego.Controller
}

func (c *CustomerController) Get() {
	o := orm.NewOrm()
	uc := models.UtCustomer{Id:123}
	err := o.Read(&uc)
	if err == nil {
		c.Data["json"] = uc
	}
	c.ServeJSON()
}


func (c *CustomerController) Post() {
	o := orm.NewOrm()
	var customer models.UtCustomer
	customer.Id = 123
	customer.OpenApiToken = "xxxxxx-xxxxs-12123"
	id, err := o.Insert(&customer)
	if err == nil {
		fmt.Println(id)
	}
}

func (c *CustomerController) NewCustomer() {
	//var requestBody map[string]interface{}
	var requestBody models.UtCustomer
	response := make(map[string]interface{})
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody); err != nil {
		response["code"] = 10001
		response["msg"] = err.Error()
	} else {
		if !models.JudgeApiIsExists(requestBody.OpenApiToken) {
			response["msg"] = "customer already exist"
			response["code"] = 10002
		} else {
			if err := models.CreateCustomer(&requestBody) ; err != nil {
				response["msg"] = err.Error()
				response["code"] = 10002
			} else {
				response["msg"] = "success"
				response["code"] = 10000
			}
		}
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}
