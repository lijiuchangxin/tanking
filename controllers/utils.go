package controllers

import (
	. "customer_managenment/api"
	model "customer_managenment/models"
	"encoding/json"
)

// 映射客户变更
func MapCustomerAlter(response *CustomerAlter, request *model.CustomerAlteration) {
	response.Id = request.Id
	response.UserId = request.UserId
	response.UserNickName = request.UserNickName
	response.AlterTime = request.AlterTime
	response.Summary = request.Summary
}

//  映射客户跟进
func MapCustomerFollow(response *CustomerFollowUp, request *model.CustomerFollowUp) {
	response.Id = request.Id
	response.UserId = request.UserId
	response.UserNickName = request.UserNickName
	response.Content = request.Content
	response.CreateAt = request.CreateAt
}

// 映射客户详情
func MapCustomerDetail(response *CustomerDetail, request *model.UtCustomer) {
	//response.Code = 0
	response.Id = request.Id
	response.CustomerNikeName = request.CustomerNikeName
	response.Desc = request.Desc
	response.Tag = request.Tag
	response.TelPhone = request.TelPhone
	response.CellPhone = request.CellPhone
	response.Email = request.Email
	response.IsVip = request.IsVip
	response.Province = request.Province
	response.City = request.City
	response.SourceChannel = request.SourceChannel
	response.CreateAt = request.CreateAt
	response.UpdatedAt = request.UpdatedAt
	response.OrganizationName = request.OrganizationName
	response.OrganizationId = request.OrganizationId
	response.OwnerGroupId = request.OwnerGroupId
	response.OwnerGroupName = request.OwnerGroupName
	response.OwnerId = request.OwnerId
	response.OwnerName = request.OwnerName
	response.TicketCount = request.TicketCount
	response.LastContactAt = request.LastContactAt
	response.LastContactImAt = request.LastContactImAt
	response.FirstContactAt = request.FirstContactAt
	response.FirstContactImAt = request.FirstContactImAt
	response.OpenApiToken = request.OpenApiToken
	alterMap := make(map[int]*CustomerAlter)
	for _, alterDetail := range request.Alters {
		alter := new(CustomerAlter)
		MapCustomerAlter(alter, alterDetail)
		alterMap[alter.Id] = alter
	}
	followMap := make(map[int]*CustomerFollowUp)
	for _, followDetail := range request.FollowUps {
		follow := new(CustomerFollowUp)
		MapCustomerFollow(follow, followDetail)
		followMap[follow.Id] = follow
	}
	response.FollowUp = followMap
	response.Alters = alterMap
}


// func GetUpdateCustomerMap(request *RequestUpdateCustomer) map[string]interface{} {将结构体映射为字典
func GetUpdateCustomerMap(request *RequestUpdateCustomer) map[string]interface{} {
	// 这个字典主要是为了最后更新值供reflect使用 对应
	nm := map[string]string{
		"tag"			:"Tag",
		"is_vip"		:"IsVip",
		"desc"			:"Desc",
		"tel_phone"		:"TelPhone",
		"cell_phone"	:"cell_phone",
	}

	// return结果
	res := make(map[string]interface{})

	// 存储 struct-- > map
	m := make(map[string]interface{})
	j, _ := json.Marshal(request)
	_ = json.Unmarshal(j, &m)
	// 用不到
	delete(m, "customer_id")
	for key, value := range m {
		if value != nil { res[nm[key]] = value }
	}
	return res

	// ??????????是个问题
	//m := make(map[string]interface{})
	//j, _ := json.Marshal(request)
	//_ = json.Unmarshal(j, &m)
	//for key, value := range m {
	//	if value != nil { m[nm[key]] = value }
	//	delete(m, key)
	//}
	//delete(m, "customer_id")
	//
	//fmt.Println(m)
	//return m
}

//func ()  {
//
//}

