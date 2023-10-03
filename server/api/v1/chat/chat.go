package api

import (
	interfaces "MyTodo/interface"
	"MyTodo/model/po/v1"
	"mime/multipart"
)

type AddChatRequest struct {
	From    uint                    `form:"from"`
	To      uint                    `form:"to"`
	Reply   string                  `form:"reply"`
	Content string                  `form:"content"`
	Files   []*multipart.FileHeader `form:"files"`
}

type AddChatResponse struct {
	interfaces.BaseResponse
}

type GetChatRequest struct {
	Page     int `json:"page" default:"1"`
	PageSize int `json:"pageSize" default:"10"`
	From     int `json:"from"`
	To       int `json:"to"`
}

type GetChatResponse struct {
	interfaces.BaseResponse
	Chats []po.Chat `json:"chats"`
}

type DelChatRequest struct {
	ID string `form:"id"`
}

type DelChatResponse struct {
	interfaces.BaseResponse
}

type ChatFriendResponse struct {
	interfaces.BaseResponse
	Friends []po.User `json:"friend"`
}
