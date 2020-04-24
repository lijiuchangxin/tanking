package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"gowebsite/apps/admin/employee/models"
)

type RoleController struct {
	beego.Controller
}

//
//func (this *RoleController) Index() {
//	this.Ctx.WriteString("hello admin's demo")
//}

// 完成 rel role operation
//参数需分情况处理
func (c *RoleController) Query() {
	resp := new(ResponseStru)
	str := c.Ctx.Input.Param(":id")
	var queryErr error
	var roles []models.UtRole
	var role models.UtRole
	if str == "" {
		roles, queryErr = models.QueryRole()
	} else {
		id, err := strconv.Atoi(str)
		if err != nil {
			logs.Error(err)
			queryErr = err
		} else {
			role, queryErr = models.QueryRoleById(id)
		}
	}

	if queryErr != nil {
		resp.RespFailSet(queryErr)
	} else {
		if str == "" {
			resp.RespSuccessSet(roles)
		} else {
			resp.RespSuccessSet(role)
		}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

//完成 add role 、rel_role_operation
func (c *RoleController) Add() {
	resp := new(ResponseStru)
	var utRole models.UtRole

	str := c.Ctx.Input.RequestBody
	err := json.Unmarshal(str, &utRole)
	if err != nil {
		logs.Error("request to object failed,", err)
		//resp.RespFailSet("requestBody2Object failed ")
	}
	num, errAdd := models.AddRole(utRole)

	if errAdd != nil || num == 0 {
		resp.RespFailSet("add failed")
	} else {
		resp.RespSuccessSet("add success")
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

//完成 update  role 、rel_role_operation
func (c *RoleController) Update() {
	resp := new(ResponseStru)
	var utRole models.UtRole
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &utRole)
	if err != nil {
		logs.Error("request to object failed,", err)
	}
	num, errUpdate := models.Update(utRole)
	if errUpdate != nil {
		resp.RespFailSet(err)
	} else if num == 0 {
		resp.RespFailSet("update : 0")
	}
	resp.RespSuccessSet("")
	c.Data["json"] = resp
	c.ServeJSON()
}

//完成 delete role 、rel_role_operation
func (c *RoleController) Delete() {
	resp := new(ResponseStru)
	var utRole models.UtRole
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &utRole)
	if err != nil {
		logs.Error("requestBody to object fail", err)
		c.Abort("")
		resp.RespFailSet(err)
	}
	num, err := models.Delete(utRole)

	if err != nil {
		logs.Error("delete is failed", err)
		resp.RespFailSet(err)
	} else if num == 0 {
		resp.RespFailSet("delete: 0")
	} else {
		resp.RespSuccessSet("delete success")
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
