package api

import (
	interfaces "MyTodo/interface"
	"MyTodo/model/po/v1"
	"mime/multipart"
	"time"

	"github.com/relvacode/iso8601"
)

type TaskCreateRequest struct {
	Topic     uint             `json:"topic"`
	Name      string           `json:"name"`
	Desc      string           `json:"desc"`
	Departure iso8601.Time     `json:"departure"`
	Arrival   iso8601.Time     `json:"arrival"`
	Cron      int64            `json:"cron"`
	Conds     []TaskCreateCond `json:"conds"`
}

type TaskCreateCond struct {
	Type  int    `json:"type"`
	Param string `json:"param"`
}

type TaskCreateResponse struct {
	interfaces.BaseResponse
}

type TaskGetRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type TaskGetResponse struct {
	interfaces.BaseResponse
	Tasks []po.DetailedTask `json:"tasks"`
}

type TaskInfoRequest struct {
	ID uint `form:"id"`
}

type TaskInfoResponse struct {
	interfaces.BaseResponse
	Name      string              `json:"name"`
	Desc      string              `json:"desc"`
	Departure time.Time           `json:"departure"`
	Arrival   time.Time           `json:"arrival"`
	Conds     []TaskInfoCondition `json:"conds"`
}

type TaskInfoCondition struct {
	Type       uint     `json:"type"`
	WantParams []string `json:"want_params"`
	GotParams  []string `json:"got_params"`
}

type TaskCommitRequest struct {
	TID   int                     `json:"tid" form:"tid"`
	Type  int                     `json:"type" form:"type"`
	Param string                  `json:"param" form:"param"`
	Files []*multipart.FileHeader `json:"files" form:"files"`
}

type TaskHasPermRequest struct {
	TID int `json:"tid" form:"tid"`
}

type TaskHasPermResposne struct {
	interfaces.BaseResponse
	IsAdmin bool `json:"has"`
}

type TaskCommitResponse struct {
	interfaces.BaseResponse
	Param string `json:"param"`
}

type TaskDeleteRequest struct{}

type TaskDeleteResponse struct{}
