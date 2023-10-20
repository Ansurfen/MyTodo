package service

import (
	api "MyTodo/api/v1/chat"
	chatController "MyTodo/controller/chat/v1"
	userController "MyTodo/controller/user/v1"
	"MyTodo/dao"
	"MyTodo/engine/v1/db"
	"MyTodo/engine/v1/starter"
	interfaces "MyTodo/interface"
	"MyTodo/model/po/v1"
	"MyTodo/utils"
	"MyTodo/utils/vfs"
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"time"
)

const (
	ChatTypeTagText = "0001"
	ChatTypeTagFile = "0002"

	ChatFileImagePath = "chat:/user/%s%s"
)

// @Summary Add Chat
// @Description add chat
// @Tags Chat
// @Accept multipart/form-data
// @Param from formData uint true "from"
// @Param to formData uint true "to"
// @Param reply formData string false "reply"
// @Param content formData string true "content"
// @Param files formData file false "files"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /chat/add [post]
func AddChat(
	ctx starter.TodoContext,
	req api.AddChatRequest) (
	interfaces.Response, error) {
	contents := []string{}
	for _, file := range req.Files {
		filename := fmt.Sprintf(ChatFileImagePath, utils.RandString(16), filepath.Ext(file.Filename))
		contents = append(contents, ChatTypeTagFile+filename)
		if err := ctx.SaveUploadFile(file, filename); err != nil {
			return ctx.ThrowWithResult(err)
		}
	}
	contents = append(contents, ChatTypeTagText+req.Content)
	chatController.Get(ctx).ChatDao.Create(&po.Chat{
		CreatedAt: time.Now(),
		From:      req.From,
		To:        req.To,
		Reply:     req.Reply,
		Content:   contents,
	})
	return api.AddChatResponse{}, nil
}

// @Summary Get Chat
// @Description get chat
// @Tags Chat
// @Param data body api.GetChatRequest true "GetChatRequest"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /chat/get [post]
func GetChat(
	ctx starter.TodoContext,
	req api.GetChatRequest) (
	interfaces.Response, error) {
	page := dao.Pagination[po.Chat]{
		Index: int64(req.Page),
		Size:  int64(req.PageSize),
	}
	err := chatController.Get(ctx).ChatDao.Find(req.From, req.To, &page)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	for i, d := range page.Data {
		for j, c := range d.Content {
			if len(c) > 4 {
				id, err := strconv.Atoi(c[0:4])
				if err != nil {
					return ctx.ThrowWithResult(err)
				}
				switch id {
				case 2:
					bucket, obejct, err := vfs.Objectf(c[4:])
					if err != nil {
						return ctx.ThrowWithResult(err)
					}
					url, err := db.OSS().Client.PresignedGetObject(context.Background(), bucket, obejct, 5*time.Hour, nil)
					if err != nil {
						return ctx.ThrowWithResult(err)
					}
					page.Data[i].Content[j] = ChatTypeTagFile + url.String()
				}
			}
		}
	}
	return api.GetChatResponse{
		Chats: page.Data,
	}, nil
}

// @Summary Del Chat
// @Description delete chat
// @Tags Chat
// @Param id formData string true "ID"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /chat/del [post]
func DelChat(
	ctx starter.TodoContext,
	req api.DelChatRequest) (
	interfaces.Response, error) {
	return api.DelChatResponse{}, chatController.Get(ctx).ChatDao.Delete(req.ID)
}

// @Summary Get friend
// @Description get friend
// @Tags Chat
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /chat/friend [get]
func Friend(
	ctx starter.TodoContext,
	req interfaces.EmptyRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.ChatFriendResponse{
		Friends: uc.DetailedFriends(),
	}, nil
}

// @Summary Get snapshot
// @Description get snapshot
// @Tags Chat
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /chat/snap [get]
func Snapshot(
	ctx starter.TodoContext,
	req interfaces.EmptyRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	fmt.Println(uc)
	return nil, nil
}
