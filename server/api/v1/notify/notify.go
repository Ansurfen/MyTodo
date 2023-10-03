package api

import (
	"MyTodo/model/po/v1"
	interfaces "MyTodo/interface"
)

type NotifyAddRequest struct {
	Title   string `json:"title" default:"This is a title"`
	Content string `json:"content" default:"This is a content"`
}

type NotifyAddResponse struct {
	interfaces.BaseResponse
}

type NotifyDelRequest struct {
	ID int `form:"id"`
}

type NotifyDelResponse struct {
	interfaces.BaseResponse
}

type NotifyGetDetailRequest struct {
	Id int `form:"id"`
}

type NotifyGetDetailResponse struct {
	interfaces.BaseResponse
	Notify po.NotifyText `json:"notify"`
}

type NotifyGetRequest interfaces.EmptyRequest

type NotifyGetResponse struct {
	interfaces.BaseResponse
	Notifications []po.NotifyText `json:"notify"`
}

type NotifyPubAddRequest struct {
	ID  int `form:"nid"`
	UID int `form:"uid"`
}

type NotifyPubAddReposne struct {
	interfaces.BaseResponse
}

type NotifyPubDelRequest struct {
	Id int `form:"nid"`
}

type NotifyPubDelResponse struct {
	interfaces.BaseResponse
}

type NotifyPubGetRequest struct{}

type NotifyPubGetResponse struct {
	interfaces.BaseResponse
	Notifications []po.NotifyPub `json:"notify"`
}

type NotifyActionAddRequest struct {
	Type  uint8  `form:"type"`
	Recv  int    `form:"id"`
	Param string `form:"param"`
}

type NotifyActionAddResponse struct {
	interfaces.BaseResponse
}

type NotifyActionDelRequest struct {
	Id int `form:"id"`
}

type NotifyActionDelResponse struct {
	interfaces.BaseResponse
}

type NotifyActionCommitRequest struct {
	Id     int `form:"id"`
	Status int `form:"status"`
}

type NotifyActionCommitResponse struct {
	interfaces.BaseResponse
}

type NotifyActionGetRequest interfaces.EmptyRequest

type NotifyActipnGetResponse struct {
	interfaces.BaseResponse
	Notifications []po.NotifyAction `json:"notify"`
}

type NotifyAllRequest struct{}

type NotifyAllResponse struct {
	interfaces.BaseResponse
	Notifications []po.Notify `json:"notify"`
}
