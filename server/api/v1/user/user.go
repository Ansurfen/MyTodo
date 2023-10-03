package api

import (
	interfaces "MyTodo/interface"
	"mime/multipart"
)

type UserSignRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email" default:"a@gmail.com"`
	Password string `json:"password" form:"password" binding:"required,gt=0" default:"root"`
}

type UserSignResposne struct {
	interfaces.BaseResponse
	JWT string `json:"jwt" form:"jwt"`
}

type UserInfoRequest interfaces.EmptyRequest

type UserInfoResponse struct {
	interfaces.BaseResponse
	Name string `json:"name"`
}

type UserGetRequest interfaces.EmptyRequest

type UserGetResponse struct {
	interfaces.BaseResponse
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

type UserEditRequest struct {
	Profile   *multipart.FileHeader `json:"profile" form:"profile"`
	Name      string                `json:"name" form:"name"`
	Email     string                `json:"email" form:"email"`
	Telephone string                `json:"telephone" form:"telephone"`
}

type UserEditResponse struct {
	interfaces.BaseResponse
}

type UserProfileRequest struct {
	Id uint `json:"id" form:"id"`
}

type UserProfileResponse struct {
	interfaces.BaseResponse
}
