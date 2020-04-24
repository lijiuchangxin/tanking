package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	. "gowebsite/apps/admin/api"
	"gowebsite/apps/admin/employee/models"
	. "gowebsite/tool/logger"
	)

type GroupController struct {
	beego.Controller
}

// 初始化组
func InitResponse(response *CommonResponse)  {
	response.Msg = "success"
	response.Code = 1
}

// 解析参数和校验参数,post
func (c *GroupController) AnalysisAndVerify(request VerifyGroupInput) bool {
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


// CreateGroup 新建组
func (c *GroupController)CreateGroup() {
	request := new(RequestCreateGroup)
	response := new(ResponseCreateGroup)
	InitResponse(&response.CommonResponse)

	if res := c.AnalysisAndVerify(request); res {
		if res, group := models.InsertNewGroup(request.GroupName); res {
			response.Code = 0
			response.Name = group.Name
			response.CreateTime = group.CreateTime
			response.Description = group.Description
			response.GroupId = group.Id
		} else {
			response.Code = 2
			response.Msg = "create group failed"
		}

	} else {
		response.Msg =  "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}


// RemoveGroup 移除组
func (c *GroupController)RemoveGroup() {
	request := new(RequestDeleteGroup)
	response := new(ResponseDeleteGroup)
	InitResponse(&response.CommonResponse)

	if c.AnalysisAndVerify(request) {
		if group := models.JudgeGroupById(request.GroupId); group != nil {
			if models.DeleteGroup(group) {
				response.Code = 0
				response.GroupId = request.GroupId
			} else {
				response.Code = 2
				response.Msg = "delete group failed"
			}
		} else {
			response.Msg = "group dose not exist"
		}
	} else {
		response.Msg =  "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}


// UpdateGroup 修改组
func (c *GroupController)UpdateGroup() {
	request := new(RequestUpdateGroup)
	response := new(ResponseUpdateGroup)
	InitResponse(&response.CommonResponse)
	if c.AnalysisAndVerify(request) {
		// 名字不重复 且 组存在
		//&& models.JudgeGroupByName(request.GroupName)
		if group := models.JudgeGroupById(request.GroupId); group != nil {
			if models.JudgeGroupExistByName(request.GroupName) {
				group.Name = request.GroupName
				if models.UpdateGroup(group) {
					response.Code = 0
					response.GroupId = request.GroupId
				} else {
					response.Code = 2
					response.Msg = "update group is failed"
				}
			} else {
				response.Msg = "group name repeat"
			}
		} else {
			response.Msg = "group dose not exist"
		}
	} else {
		response.Msg =  "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}


// GetGroupDetail 获取组详细信息
func (c *GroupController)GetGroupDetail() {
	request := new(RequestGroupDetail)
	response := new(ResponseGroupDetail)
	InitResponse(&response.CommonResponse)
	if res, err := c.GetInt("group_id"); err == nil {
		request.GroupId = res
		if request.VerifyInputPara() {
			if group := models.JudgeGroupById(request.GroupId); group != nil {
				if res := models.GetGroupDetail(group); res != nil {
					response.Code = 0
					MapGroupDetail(&response.GroupDetail, group)
				} else {
					response.Code = 2
					response.Msg = "get group detail failed"
				}
			} else {
				response.Msg = "group dose not exist"
			}
		} else {
			response.Msg =  "incoming parameter error"
		}
	} else {
		response.Msg =  "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}

// GroupListByPage
func (c *GroupController)GroupListByPage()  {
	request := new(RequestGroupList)
	response := new(ResponseGroupList)
	InitResponse(&response.CommonResponse)

	request.CurrPage, _ = c.GetInt("curr_page")
	request.PageSize, _ = c.GetInt("page_size")

	// 校验入参是否正确
	//if res := c.AnalysisAndVerify(request); res {
	if res := request.VerifyInputPara(); res {
		if res, count, err := models.GroupListPageOut(request.CurrPage, request.PageSize); err != nil {
			response.Code = 2
			if count == 0 {
				response.Msg = "error in getting total number of group details"
			} else {
				response.Msg = "failed to get group details on page"
			}
		} else {
			m := make(map[int]*GroupDetail)
			response.Groups = m
			response.Code = 0
			for num, group := range res {
				groupDetail := new(GroupDetail)
				MapGroupDetail(groupDetail, group)
				response.Groups[num] = groupDetail
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

// GroupAddAgent 组添加员工
func (c *GroupController)GroupAddAgent()  {
	request := new(RequestGroupOperateAgent)
	response := new(ResponseGroupOperateAgent)
	InitResponse(&response.CommonResponse)

	if res := c.AnalysisAndVerify(request); res {
		// 判断组是否存在
		if group := models.JudgeGroupById(request.GroupId, ); group != nil {
			if models.AddAgent2Group(group, request.AgentId) {
				response.GroupId = request.GroupId
				response.AgentId = request.AgentId
				response.Code = 0
			} else {
				response.Code = 2
				response.Msg = "add agent failed"
			}
		} else {
			response.Msg = "add agent failed, because group dose not exist"
		}

	} else {
		response.Msg = "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
	return
}


// GroupRemoveAgent 组删除员工
func (c *GroupController)GroupRemoveAgent() {
	request := new(RequestGroupOperateAgent)
	response := new(ResponseGroupOperateAgent)
	InitResponse(&response.CommonResponse)

	if res := c.AnalysisAndVerify(request); res {
		// 判断组是否存在
		if group := models.JudgeGroupById(request.GroupId, ); group != nil {
			if models.RemoveAgentFromGroup(request.GroupId, request.AgentId) {
				response.GroupId = request.GroupId
				response.AgentId = request.AgentId
				response.Code = 0
			} else {
				response.Code = 2
				response.Msg = "remove agent failed"
			}
		} else {
			response.Msg = "remove agent failed, because group dose not exist"
		}

	} else {
		response.Msg = "incoming parameter error"
	}
	c.Data["json"] = response
	c.ServeJSON()
}



// TODO 换个位置
// 映射组详情
func MapGroupDetail(response *GroupDetail, request *models.UtEmployeeGroup) {
	//response.Code = 0
	response.GroupId = request.Id
	response.Name = request.Name
	response.Description = request.Description
	response.CreateTime = request.CreateTime
	agentMap := make(map[int]*models.UtEmployee)
	for _, agent := range request.UtEmployee {
		agentMap[agent.Id] = agent
	}
	response.Employees = agentMap
}
