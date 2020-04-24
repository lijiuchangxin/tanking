package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	. "gowebsite/apps/admin/employee/models"
)

type EmployeeController struct {
	beego.Controller
}

func (e *EmployeeController) Query() {

	resp := new(ResponseStru)

	result, err := QueryEmployee()
	if err != nil {
		logs.Error(err)
	}
	resp.RespSuccessSet(result)

	e.Data["json"] = resp
	e.ServeJSON()

}

func (e *EmployeeController) Add() {

	str := e.Ctx.Input.RequestBody
	var utEmployee UtEmployee
	resp := new(ResponseStru)
	err := json.Unmarshal(str, &utEmployee)
	if err != nil {
		logs.Error(err)
	}
	err = AddEmployee(utEmployee)
	if err != nil {
		resp.RespFailSet(err)
	} else {
		resp.RespSuccessSet("ok")
	}
	e.Data["json"] = resp
	e.ServeJSON()

}
