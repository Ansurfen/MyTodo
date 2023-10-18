package service

import (
	api "MyTodo/api/v1/task"
	taskController "MyTodo/controller/task/v1"
	userController "MyTodo/controller/user/v1"
	"MyTodo/engine/v1/db"
	"MyTodo/engine/v1/starter"
	interfaces "MyTodo/interface"
	"MyTodo/model/po/v1"
	"MyTodo/utils"
	"fmt"
	"path/filepath"
	"strings"
)

// @Summary Create Task
// @Description create task
// @Tags Task
// @Accept json
// @Param data body api.TaskCreateRequest true "TaskCreateRequest"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /task/add [post]
func CreateTask(ctx starter.TodoContext, req api.TaskCreateRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(utils.ErrUserNotFound)
	}
	u := uc.User
	task := po.Task{
		TTID:      req.Topic,
		User:      u.ID,
		Name:      req.Name,
		Desc:      req.Desc,
		Departure: req.Departure.Time,
		Arrival:   req.Arrival.Time,
	}
	tc := taskController.Get(ctx)
	if err := tc.CreateTask(&task); err != nil {
		return ctx.ThrowWithResult(err)
	}
	for _, cond := range req.Conds {
		if cond.Type == po.COND_QR {
			cond.Param = utils.RandString(8)
		}
		err := tc.CreateBound(int(task.ID), cond.Type, cond.Param)
		if err != nil {
			return ctx.ThrowWithResult(err)
		}
	}
	return api.TaskCreateResponse{}, nil
}

func DeleteTask(req api.TaskDeleteRequest) (api.TaskDeleteResponse, error) {
	return api.TaskDeleteResponse{}, nil
}

func InfoTask(ctx starter.TodoContext, req api.TaskInfoRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return api.TaskCommitResponse{}, err
	}
	return taskController.Get(ctx).InfoTask(uc.User.ID, req.ID)
}

// @Summary Get Task
// @Description Get task
// @Tags Task
// @Param page formData int true "page"
// @Param limit formData int true "limit"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /task/get [get]
func GetTask(ctx starter.TodoContext, req api.TaskGetRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return api.TaskCommitResponse{}, err
	}
	res, err := taskController.Get(ctx).GetTaskByUID(uc.User.ID)
	if err != nil {
		return api.TaskGetResponse{}, err
	}
	return api.TaskGetResponse{
		Tasks: res,
	}, nil
}

// @Summary Commit Task
// @Description Commit task
// @Tags Task
// @Accept multipart/form-data
// @Param tid formData int true "tid"
// @Param type formData int true "type"
// @Param param formData string true "param"
// @Param files formData file true "files"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /task/commit [post]
func CommitTask(
	ctx starter.TodoContext,
	req api.TaskCommitRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return api.TaskCommitResponse{}, err
	}
	_, err = taskController.Get(ctx).GetAllCond(req.TID, req.Type)
	if err != nil {
		return api.TaskCommitResponse{}, err
	}
	// validates whether the condition matches
	param := ""
	switch req.Type {
	case po.COND_CONTENT:
		taskController.Get(ctx).CreateCommit(int(uc.User.ID), req.TID, req.Type, req.Param)
	case po.COND_FILE:
		files := []string{}
		for _, file := range req.Files {
			filename := utils.RandString(8)
			err = ctx.Context().SaveUploadedFile(file,
				filepath.Join("./assets/task/file", filename+filepath.Ext(file.Filename)))
			if err != nil {
				return api.TaskCommitResponse{}, err
			}
			files = append(files, filename)
		}
		err = taskController.Get(ctx).CreateCommit(int(uc.User.ID), req.TID, req.Type, strings.Join(files, ","))
		if err != nil {
			return ctx.ThrowWithResult(err)
		}
	case po.COND_IMAGE:
		files := []string{}
		for _, img := range req.Files {
			filename := utils.RandString(8)
			err = ctx.Context().SaveUploadedFile(img, fmt.Sprintf(TaskFilePath, filename))
			if err != nil {
				return api.TaskCommitResponse{}, err
			}
			files = append(files, filename)
		}
		err = taskController.Get(ctx).CreateCommit(int(uc.User.ID), req.TID, req.Type, strings.Join(files, ","))
		if err != nil {
			return api.TaskCommitResponse{}, err
		}
	case po.COND_TIMER:

	case po.COND_HAND:
		err = taskController.Get(ctx).CreateCommit(int(uc.User.ID), req.TID, req.Type, "")
		if err != nil {
			return api.TaskCommitResponse{}, err
		}
	case po.COND_QR:
		res := strings.Split(req.Param, ",")
		key := fmt.Sprintf("qr_%s", res[0])
		taskController.Get(ctx).CreateCommit(int(uc.User.ID), req.TID, req.Type, "")
		if db.Redis().Get(key) == res[1] {
			return api.TaskCommitResponse{}, nil
		}
		return api.TaskCommitResponse{
			Param: "QR timeout",
		}, nil
	case po.COND_LOCATE:
		res := strings.SplitN(req.Param, "---", 2)
		if len(res) == 2 {
			id := utils.RandString(8)
			err := utils.Base64ToFile(res[1][len("data:image/png;base64,"):], fmt.Sprintf(TaskPath, id))
			if err != nil {
				return api.TaskCommitResponse{}, err
			}
			err = taskController.Get(ctx).CreateCommit(int(uc.User.ID), req.TID, req.Type, res[0]+","+id)
			if err != nil {
				return api.TaskCommitResponse{}, err
			}
			param = id
		}
	}
	return api.TaskCommitResponse{
		Param: param,
	}, nil
}

const SQL_HasPerm = `SELECT EXISTS (
    SELECT 1
    FROM task t
    JOIN topic tp ON t.tt_id = tp.id
    WHERE tp.creator = %d AND t.id = %d
) AS is_admin;`

// @Summary Has perm Task
// @Description Has perm task
// @Tags Task
// @Accept multipart/form-data
// @Param tid formData int true "tid"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /task/perm_check [post]
func TaskHasPerm(
	ctx starter.TodoContext,
	req api.TaskHasPermRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	res := api.TaskHasPermResposne{}
	err = db.SQL().Table("task").
		Select("EXISTS (SELECT 1 FROM task t JOIN topic tp ON t.tt_id = tp.id WHERE tp.creator = ? AND t.id = ?) AS is_admin", uc.User.ID, req.TID).
		Scan(&res).Error
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return res, nil
	// return res, db.SQL().RawExec(SQL_HasPerm, uc.User.ID, req.TID).Scan(&res).Error
}

func StatisticTask(ctx starter.TodoContext) {}

func CommitListTask(ctx starter.TodoContext) {}

const (
	TaskFilePath = "./assets/task/file/%s.png"
	TaskPath     = "./assets/task/%s.png"
)

// @Summary Get task image
// @Description get task image
// @Tags Task
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /task/image/{id} [get]
func TaskImage(ctx starter.TodoContext) {
	id := ctx.Context().Param("id")
	if len(id) > 0 {
		ctx.Context().File(fmt.Sprintf(TaskFilePath, id))
	}
}

// @Summary Get task image of location
// @Description get task image of location
// @Tags Task
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /task/locate/{id} [get]
func TaskLocateImage(ctx starter.TodoContext) {
	id := ctx.Context().Param("id")
	if len(id) > 0 {
		ctx.Context().File(fmt.Sprintf(TaskPath, id))
	}
}
