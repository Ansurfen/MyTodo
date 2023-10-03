package api

import (
	interfaces "MyTodo/interface"
	"MyTodo/model/po/v1"
)

type CreateTopicRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CreateTopicResponse struct {
	interfaces.BaseResponse
	InviteCode string `json:"invite_code"`
}

type SubscribeTopicRequest struct {
	InviteCode string `json:"invite_code"`
}

type SubscribeTopicResponse struct {
	interfaces.BaseResponse
}

type GetSubscribedTopicRequest interfaces.EmptyRequest

type GetSubscribedTopicResponse struct {
	interfaces.BaseResponse
	Topics []po.SubscribedTopic `json:"topics"`
}

type JoinTopicRequest struct {
	InviteCode string `json:"invite_code" form:"invite_code"`
}

type JoinTopicResponse struct {
	interfaces.BaseResponse
}

type GetTopicSubsriberRequest struct {
	TTID uint `form:"tt_id"`
}

type GetTopicSubsriberResponse struct {
	interfaces.BaseResponse
}

type DeleteTopicRequset struct {
	TTID uint `form:"id"`
}

type DeleteTopicResponse struct {
	interfaces.BaseResponse
}

type GrantAdminTopicRequest struct {
	TTID uint `form:"id"`
}

type GrantAdminTopicResponse struct {
	interfaces.BaseResponse
}

type UngrantAdminTopicRequest struct {
	TTID uint `form:"tt_id"`
}

type UngrantAdminTopicResponse struct {
	interfaces.BaseResponse
}

type MigrateTopicAdminRequest struct {
	TTID     uint `form:"tt_id"`
	NewAdmin uint `form:"uid"`
}

type MigrateTopicAdminResponse struct {
	interfaces.BaseResponse
}

type HistoryTopicRequest interfaces.EmptyRequest

type HisotryTopicResponse struct {
	interfaces.BaseResponse
	Topics []*po.Topic `json:"topic"`
}

type RemoveTopicMemberRequest struct {
	TTID uint `form:"tt_id"`
	User int  `form:"uid"`
}

type RemoveTopicMemberResponse struct {
	interfaces.BaseResponse
}

type ExitTopicRequest struct {
	TTID uint `form:"tt_id"`
}

type ExitTopicResponse struct {
	interfaces.BaseResponse
}

type ImportTopicMemeberRequest interfaces.EmptyRequest

type ImportTopicMemeberResponse struct {
	interfaces.BaseResponse
}

type GetTopicMemberRequest struct {
	ID uint `json:"id" form:"id"`
}

type GetTopicMemberResponse struct {
	interfaces.BaseResponse
	Members []po.SubscribedMemeber `json:"member"`
}
