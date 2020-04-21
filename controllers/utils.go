package controllers

import (
	. "customer_managenment/api"
	model "customer_managenment/models"
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
func MapCustomerDetail(response *ResponseNewCustomer, request *model.UtCustomer) {
	response.Code = 0
	response.UtCustomer.Id = request.Id
	response.UtCustomer.CustomerNikeName = request.CustomerNikeName
	response.UtCustomer.Desc = request.Desc
	response.UtCustomer.Tag = request.Tag
	response.UtCustomer.TelPhone = request.TelPhone
	response.UtCustomer.CellPhone = request.CellPhone
	response.UtCustomer.Email = request.Email
	response.UtCustomer.IsVip = request.IsVip
	response.UtCustomer.Province = request.Province
	response.UtCustomer.City = request.City
	response.UtCustomer.SourceChannel = request.SourceChannel
	response.UtCustomer.CreateAt = request.CreateAt
	response.UtCustomer.UpdatedAt = request.UpdatedAt
	response.UtCustomer.OrganizationName = request.OrganizationName
	response.UtCustomer.OrganizationId = request.OrganizationId
	response.UtCustomer.OwnerGroupId = request.OwnerGroupId
	response.UtCustomer.OwnerGroupName = request.OwnerGroupName
	response.UtCustomer.OwnerId = request.OwnerId
	response.UtCustomer.OwnerName = request.OwnerName
	response.UtCustomer.TicketCount = request.TicketCount
	response.UtCustomer.LastContactAt = request.LastContactAt
	response.UtCustomer.LastContactImAt = request.LastContactImAt
	response.UtCustomer.FirstContactAt = request.FirstContactAt
	response.UtCustomer.FirstContactImAt = request.FirstContactImAt
	response.UtCustomer.OpenApiToken = request.OpenApiToken
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
	response.UtCustomer.FollowUp = followMap
	response.UtCustomer.Alters = alterMap
}