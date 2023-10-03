package api

import (
	interfaces "MyTodo/interface"
	"MyTodo/model/po/v1"
	"mime/multipart"
)

type CreatePostRequest struct {
	Content string                  `json:"content" form:"content"`
	Images  []*multipart.FileHeader `form:"files"`
}

type CreatePostResponse struct {
	interfaces.BaseResponse
}

type DeletePostRequest struct {
	Pid int
}

type DeletePostResponse struct {
	interfaces.BaseResponse
}

type EditPostRequest struct {
	PID     uint      `json:"pid" form:"pid"`
	UID     uint      `json:"uid" form:"uid"`
	Content string    `json:"content" form:"content"`
	Images  po.Images `json:"count" form:"count"`
}

type EditPostResponse struct {
	interfaces.BaseResponse
}

type GetPostRequest struct {
	Page  int `json:"page" form:"page"`
	Count int `json:"count" form:"count"`
}

type GetPostResponse struct {
	interfaces.BaseResponse
	Posts []po.SnapshotPost `json:"posts"`
}

type InfoPostRequest struct{}

type InfoPostResponse struct {
	interfaces.BaseResponse
}

type FavoritePostRequest struct {
	Pid uint `json:"pid" form:"pid"`
}

type FavoritePostResponse struct {
	interfaces.BaseResponse
}

type UnfavoritePostRequest struct {
	Pid int `json:"pid" form:"pid"`
}

type UnfavoritePostResponse struct {
	interfaces.BaseResponse
}

type GetPostFavoriteRequest struct {
	Pid int `json:"pid" form:"pid"`
}

type GetPostFavoriteResponse struct {
	interfaces.BaseResponse
	Count int64 `json:"count"`
}

type GetCommentPostRequest struct {
	PID      int `form:"pid"`
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type GetCommentPostResponse struct {
	interfaces.BaseResponse
	Comments []po.DetailedPostComment `json:"comments"`
}

type CreateCommentPostRequest struct {
	PID     int                     `form:"pid"`
	Reply   uint                    `form:"reply"`
	Content string                  `form:"content"`
	Images  []*multipart.FileHeader `form:"files"`
}

type CreateCommentPostResponse struct {
	interfaces.BaseResponse
}

type EditCommentPostRequest struct {
	ID      string                  `form:"id"`
	Content string                  `form:"content"`
	Images  []*multipart.FileHeader `form:"files"`
}

type EditCommentPostResponse struct {
	interfaces.BaseResponse
}

type DelCommentPostRequest struct {
	Id string `form:"id"`
}

type DelCommentPostResponse struct {
	interfaces.BaseResponse
}

type PostCommentReplyCreateRequest struct {
	Id      string                  `form:"id"`
	Reply   uint                    `form:"reply"`
	Content string                  `form:"content"`
	Images  []*multipart.FileHeader `form:"files"`
}

type PostCommentReplyCreateResponse struct {
	interfaces.BaseResponse
}

type PostCommentReplyEditRequest struct {
	CommentID string                  `form:"comment_id"`
	ReplyID   string                  `form:"reply_id"`
	Content   string                  `form:"content"`
	Images    []*multipart.FileHeader `form:"files"`
}

type PostCommentReplyEditResponse struct {
	interfaces.BaseResponse
}

type PostCommentReplyDeleteRequest struct {
	CommentID string `form:"comment_id"`
	ReplyID   string `form:"reply_id"`
}

type PostCommentReplyDeleteResponse struct {
	interfaces.BaseResponse
}

type PostCommentFavoriteRequest struct {
	CommentID string `form:"comment_id"`
}

type PostCommentUnfavoriteRequest struct {
	CommentID string `form:"comment_id"`
}

type PostCommentFavoriteCountRequest struct {
	CommentID string `form:"comment_id"`
}

type PostCommentFavoriteCountResponse struct {
	interfaces.BaseResponse
	Count int64 `json:"count"`
}

type PostDetailResponse struct {
	interfaces.BaseResponse
	po.DetailedPost
}
