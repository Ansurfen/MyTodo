package service

import (
	api "MyTodo/api/v1/notify"
	notifyController "MyTodo/controller/notify/v1"
	topicController "MyTodo/controller/topic/v1"
	userController "MyTodo/controller/user/v1"
	"MyTodo/engine/v1/starter"

	interfaces "MyTodo/interface"
	"MyTodo/model/po/v1"
	"MyTodo/utils"
	"fmt"
	"strconv"
)

// @Summary Notify create
// @Description create notify
// @Tags Notify
// @Accept json
// @Param data body api.NotifyAddRequest true "NotifyAddRequest"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/add [post]
func NotifyAdd(ctx starter.TodoContext, req api.NotifyAddRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	err = nc.NotifyDao.Create(&po.NotifyText{
		Creator: int(uc.User.ID),
		Content: req.Content,
	})
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.NotifyAddResponse{}, nil
}

// @Summary Notify delete
// @Description delete notify
// @Tags Notify
// @Accept multipart/form-data
// @Param id formData int true "ID"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/del [post]
func NotifyDel(ctx starter.TodoContext, req api.NotifyDelRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	err = nc.NotifyDao.Delete(int(uc.User.ID), req.ID)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.NotifyDelResponse{}, nil
}

// @Summary Notify get
// @Description get notify
// @Tags Notify
// @Accept multipart/form-data
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/get [get]
func NotifyGet(ctx starter.TodoContext, req api.NotifyGetRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	res, err := nc.NotifyDao.FindByCreator(int(uc.User.ID))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.NotifyGetResponse{
		Notifications: res,
	}, nil
}

// @Summary Notify get detail
// @Description get notify detail
// @Tags Notify
// @Accept multipart/form-data
// @Param id formData string true "id"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/detail [get]
func NotifyGetDetial(
	ctx starter.TodoContext,
	req api.NotifyGetDetailRequest) (
	interfaces.Response, error) {
	nc := notifyController.Get(ctx)
	var res api.NotifyGetDetailResponse
	err := nc.NotifyDao.UnscopedFindOneByID(req.Id, &res.Notify)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return res, nil
}

// @Summary Notify pub create
// @Description create notify pub
// @Tags Notify
// @Accept multipart/form-data
// @Param nid formData int true "nid"
// @Param uid formData int true "UID"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/pub/add [post]
func NotifyPubAdd(ctx starter.TodoContext, req api.NotifyPubAddRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	notify, err := nc.NotifyDao.FindOneByID(req.ID)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	if notify.Creator != int(uc.User.ID) {
		return ctx.ThrowWithResult(utils.ErrPermDenied)
	}
	err = nc.NotifyPubDao.Create(&po.NotifyPub{
		NID: uint(req.ID),
		UID: uint(req.UID),
	})
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.NotifyAddResponse{}, nil
}

// @Summary Notify pub delete
// @Description delete notify pub
// @Tags Notify
// @Accept multipart/form-data
// @Param nid formData int true "nid"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/pub/del [post]
func NotfiyPubDel(ctx starter.TodoContext, req api.NotifyPubDelRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)

	pub, err := nc.NotifyPubDao.FindByID(req.Id)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}

	notify, err := nc.NotifyDao.FindOneByID(int(pub.ID))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	if notify.ID != uc.User.ID {
		return ctx.ThrowWithResult(utils.ErrPermDenied)
	}
	err = nc.NotifyPubDao.Delete(req.Id)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.NotifyDelResponse{}, nil
}

// @Summary Notify pub get
// @Description get notify pub
// @Tags Notify
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/pub/get [get]
func NotifyPubGet(ctx starter.TodoContext, req api.NotifyGetRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	res, err := nc.NotifyPubDao.FindByCreator(int(uc.User.ID))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.NotifyPubGetResponse{
		Notifications: res,
	}, nil
}

// @Summary Notify action create
// @Description type: 1 (add friend), 2 (invite friend)
// @Tags Notify
// @Param type formData int true "type"
// @Param id formData int true "recv"
// @Param param formData string false "param"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/action/add [post]
func NotifyActionAdd(
	ctx starter.TodoContext,
	req api.NotifyActionAddRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	state := po.A_StateUnknown
	switch req.Type {
	case po.A_NotifyAddFriend:
		state = po.A_StateWait
	case po.A_NotifyInviteFriend:
		state = po.A_StateWait
	default:
		return ctx.ThrowWithResult(fmt.Errorf("invalid state"))
	}
	nc.NotifyActionDao.Create(&po.NotifyAction{
		Sender:   uc.User.ID,
		Receiver: uint(req.Recv),
		Type:     req.Type,
		Status:   uint8(state),
		Param:    req.Param,
	})
	return api.NotifyActionAddResponse{}, nil
}

// @Summary Notify action delete
// @Description delete notify action
// @Tags Notify
// @Param type formData int true "type"
// @Param id formData int true "recv"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/action/del [post]
func NotifyActionDel(ctx starter.TodoContext,
	req api.NotifyActionDelRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	err = nc.NotifyActionDao.Delete(int(uc.User.ID), req.Id)
	if err != nil {
		return api.NotifyActionDelResponse{}, err
	}
	return api.NotifyActionDelResponse{}, nil
}

// @Summary Notify action get
// @Description get notify action
// @Tags Notify
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/action/get [get]
func NotifyActionGet(
	ctx starter.TodoContext,
	req api.NotifyActionGetRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	res, err := nc.NotifyActionDao.FindAllByUID(int(uc.User.ID))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.NotifyActipnGetResponse{
		Notifications: res,
	}, nil
}

// @Summary Notify action commit
// @Description status: 2 (confirm), 3 (refuse)
// @Tags Notify
// @Param status formData int true "status"
// @Param id formData int true "id"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/action/commit [post]
func NotifyActionCommit(
	ctx starter.TodoContext,
	req api.NotifyActionCommitRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	act, err := nc.NotifyActionDao.FindOneByID(req.Id)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	switch act.Type {
	case po.A_NotifyAddFriend:
		if req.Status == po.A_StateConfirm {
			err = uc.Follow(int(act.Sender))
			if err != nil {
				return ctx.ThrowWithResult(err)
			}
		}
	case po.A_NotifyInviteFriend:
		if req.Status == po.A_StateConfirm {
			tc := topicController.Get(ctx)
			id, err := strconv.Atoi(act.Param)
			if err != nil {
				return ctx.ThrowWithResult(err)
			}
			err = tc.SubscribeDao.Create(&po.Subscribe{
				UID:  act.Receiver,
				TTID: uint(id),
				Perm: po.PermMember,
			})
			if err != nil {
				return ctx.ThrowWithResult(err)
			}
		}
	}
	err = nc.NotifyActionDao.Commit(req.Id, int(uc.User.ID), req.Status)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.NotifyActionCommitResponse{}, nil
}

// @Summary Notify all
// @Description all notify
// @Tags Notify
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /notify/all [get]
func NotifyAll(ctx starter.TodoContext, req api.NotifyAllRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	nc := notifyController.Get(ctx)
	notifies := []po.Notify{}
	res, err := nc.NotifyActionDao.FindAllByUID(int(uc.User.ID))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	for _, n := range res {
		notifies = append(notifies, po.Notify{
			Id:       int(n.ID),
			Type:     n.Type,
			Status:   n.Status,
			CreateAt: n.CreatedAt,
			Param:    strconv.Itoa(int(n.Sender)),
		})
	}
	res2, err := nc.NotifyPubDao.FindByCreator(int(uc.User.ID))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	for _, n := range res2 {
		notifies = append(notifies, po.Notify{
			Id:       int(n.ID),
			Type:     po.A_NotifyText,
			Status:   po.A_StateUnknown,
			CreateAt: n.CreatedAt,
		})
	}
	return api.NotifyAllResponse{
		Notifications: notifies,
	}, nil
}
