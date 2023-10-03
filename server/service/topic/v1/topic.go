package service

import (
	api "MyTodo/api/v1/topic"
	topicController "MyTodo/controller/topic/v1"
	userController "MyTodo/controller/user/v1"
	"MyTodo/engine/v1/starter"
	interfaces "MyTodo/interface"
	"MyTodo/model/po/v1"
	"MyTodo/utils"
	"fmt"
	"strconv"
)

// @Summary Create Topic
// @Description create Topic
// @Tags Topic
// @Accept multipart/form-data
// @Param name formData string true "Name"
// @Param desc formData string true "description"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/add [post]
func CreateTopic(
	ctx starter.TodoContext,
	req api.CreateTopicRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return api.CreateTopicResponse{}, utils.ErrUserNotFound
	}
	tc := topicController.Get(ctx)
	code, err := tc.CreateTopic(uc.User.ID, req.Name, req.Desc)
	if err != nil {
		return api.CreateTopicResponse{}, err
	}
	return api.CreateTopicResponse{
		InviteCode: code,
	}, nil
}

// @Summary Delete Topic
// @Description delete Topic
// @Tags Topic
// @Accept multipart/form-data
// @Param id formData int true "task topic id"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/del [post]
func DeleteTopic(
	ctx starter.TodoContext,
	req api.DeleteTopicRequset) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return api.DeleteTopicResponse{}, err
	}
	tc := topicController.Get(ctx)
	if !tc.IsAdmin(int(req.TTID), int(uc.User.ID)) {
		return ctx.ThrowWithResult(utils.ErrPermDenied)
	}
	err = tc.TopicDao.Delete(req.TTID)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.DeleteTopicResponse{}, nil
}

// @Summary Get history of Topic
// @Description get history of Topic
// @Tags Topic
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/history [get]
func HistoryTopic(
	ctx starter.TodoContext,
	req api.HistoryTopicRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	tc := topicController.Get(ctx)
	res, err := tc.TopicDao.FindAll(int(uc.User.ID))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.HisotryTopicResponse{
		Topics: res,
	}, nil
}

// @Summary Remove The member of Topic
// @Description Remove The member of Topic
// @Tags Topic
// @Accept multipart/form-data
// @Param id formData int true "task topic id"
// @Param uid formData int true "user id"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/member/rm [post]
func RemoveTopicMember(
	ctx starter.TodoContext,
	req api.RemoveTopicMemberRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	if uc.User.ID == uint(req.User) {
		return ctx.ThrowWithResult(utils.ErrPermDenied)
	}
	tc := topicController.Get(ctx)
	if !tc.IsManager(int(req.TTID), int(uc.User.ID)) {
		return ctx.ThrowWithResult(utils.ErrPermDenied)
	}
	err = tc.SubscribeDao.Delete(uc.User.ID, req.TTID)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.RemoveTopicMemberResponse{}, nil
}

// @Summary Exits Topic
// @Description exits Topic
// @Tags Topic
// @Accept multipart/form-data
// @Param id formData int true "task topic id"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/exit [post]
func ExitTopic(
	ctx starter.TodoContext,
	req api.ExitTopicRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	tc := topicController.Get(ctx)
	if !tc.IsMemberOrManager(int(req.TTID), int(uc.User.ID)) {
		return ctx.ThrowWithResult(fmt.Errorf("try to migrate or delete topic"))
	}
	err = tc.SubscribeDao.Delete(uc.User.ID, req.TTID)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.ExitTopicResponse{}, nil
}

func ImportTopicMemeber(
	ctx starter.TodoContext,
	req api.ImportTopicMemeberRequest) (
	interfaces.Response, error) {
	return api.ImportTopicMemeberResponse{}, nil
}

// @Summary Migrate perm Topic
// @Description migrate Topic
// @Tags Topic
// @Accept multipart/form-data
// @Param id formData int true "task topic id"
// @Param uid formData int true "user id"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/migrate [post]
func MigrateTopicAdmin(
	ctx starter.TodoContext,
	req api.MigrateTopicAdminRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	tc := topicController.Get(ctx)
	if !tc.IsAdmin(int(req.TTID), int(uc.User.ID)) {
		return ctx.ThrowWithResult(utils.ErrPermDenied)
	}
	if !tc.IsJoin(int(req.TTID), int(req.NewAdmin)) {
		return ctx.ThrowWithResult(fmt.Errorf("user not found"))
	}
	err = tc.SubscribeDao.UpdatePerm(int(uc.User.ID), int(req.TTID), po.PermMember)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	err = tc.SubscribeDao.UpdatePerm(int(req.NewAdmin), int(req.TTID), po.PermAdmin)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.MigrateTopicAdminResponse{}, nil
}

// @Summary Grant perm Topic
// @Description grant Topic
// @Tags Topic
// @Accept multipart/form-data
// @Param id formData int true "task topic id"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/grant [post]
func GrantAdminTopic(
	ctx starter.TodoContext,
	req api.GrantAdminTopicRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	tc := topicController.Get(ctx)
	if !tc.IsAdmin(int(req.TTID), int(uc.User.ID)) {
		return ctx.ThrowWithResult(utils.ErrPermDenied)
	}
	err = tc.SubscribeDao.UpdatePerm(int(uc.User.ID), int(req.TTID), po.PermManager)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.GrantAdminTopicResponse{}, nil
}

// @Summary Ungrant perm Topic
// @Description ungrant Topic
// @Tags Topic
// @Accept multipart/form-data
// @Param tt_id formData int true "task topic id"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/ungrant [post]
func UngrantAdminTopic(
	ctx starter.TodoContext,
	req api.UngrantAdminTopicRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	tc := topicController.Get(ctx)
	if !tc.IsAdmin(int(req.TTID), int(uc.User.ID)) {
		return ctx.ThrowWithResult(utils.ErrPermDenied)
	}
	err = tc.SubscribeDao.UpdatePerm(int(uc.User.ID), int(req.TTID), po.PermMember)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.UngrantAdminTopicResponse{}, nil
}

// @Summary Get Topic
// @Description get Topic
// @Tags Topic
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/get [get]
func GetSubscribedTopic(
	ctx starter.TodoContext,
	req api.GetSubscribedTopicRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	tc := topicController.Get(ctx)
	res, err := tc.GetSubscribedTopic(int(uc.User.ID))
	return api.GetSubscribedTopicResponse{
		Topics: res,
	}, err
}

// @Summary Subscribes topic
// @Description Subscribes topic
// @Tags Topic
// @Accept json
// @Param data body api.JoinTopicRequest true "JoinTopicRequest"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/sub [post]
func SubscirbeTopic(
	ctx starter.TodoContext,
	req api.JoinTopicRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	tc := topicController.Get(ctx)
	return api.JoinTopicResponse{}, tc.SubscribeTopic(uc.User.ID, req.InviteCode)
}

// @Summary Get members of topic
// @Description get members of topic
// @Tags Topic
// @Param tt_id path int true "task topic id"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /topic/member/{id} [get]
func GetTopicMember(
	ctx starter.TodoContext,
	req api.GetTopicMemberRequest) (
	interfaces.Response, error) {
	tc := topicController.Get(ctx)
	id := ctx.Context().Param("id")
	i, _ := strconv.Atoi(id)
	m, err := tc.GetSubscribedMemeber(uint(i))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.GetTopicMemberResponse{
		Members: m,
	}, nil
}
