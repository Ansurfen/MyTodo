package controller

import (
	api "MyTodo/api/v1/task"
	"MyTodo/dao"
	"MyTodo/engine/v1/db"
	"MyTodo/engine/v1/starter"
	"MyTodo/model/po/v1"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	ctx           *gin.Context
	TaskDao       dao.TaskDao
	TaskCondDao   dao.TaskCondDao
	TaskBoundDao  dao.TaskBoundDao
	TaskCommitDao dao.TaskCommitDao
}

func Get(ctx starter.TodoContext) *TaskController {
	return &TaskController{ctx: ctx.Context()}
}

func (c *TaskController) CreateTask(task *po.Task) error {
	return c.TaskDao.Create(task)
}

func (c *TaskController) CreateCond(_type int, param string) error {
	c.TaskCondDao.Create(po.TaskCond{})
	return nil
}

func (c *TaskController) CreateBound(tid, _type int, param string) error {
	return c.TaskBoundDao.Create(po.TaskBound{
		TID:   uint(tid),
		TCID:  uint(_type),
		Param: param,
	})
}

func (c *TaskController) CreateCommit(uid, tid, tcid int, param string) error {
	return c.TaskCommitDao.Create(po.TaskCommit{
		UID:   uint(uid),
		TID:   uint(tid),
		TCID:  uint(tcid),
		Param: param,
	})
}

func (c *TaskController) GetTask(page, limit int) ([]po.Task, error) {
	if limit < 0 {
		limit = 10
	}
	if page < 0 {
		page = 1
	}
	return c.TaskDao.Find(page, limit)
}

func (c *TaskController) GetCond(tid int) {

}

func (c *TaskController) GetAllCond(tid, tcid int) (ret []po.TaskBound, err error) {
	return c.TaskBoundDao.FindMany(po.TaskBound{
		TID:  uint(tid),
		TCID: uint(tcid),
	})
}

func (c *TaskController) InfoTask(uid, tid uint) (api.TaskInfoResponse, error) {
	task, err := c.TaskDao.FindById(int(tid))
	if err != nil {
		return api.TaskInfoResponse{}, err
	}
	commits := []po.TaskCommit{}
	bounds := []po.TaskBound{}
	db.SQL().Raw(fmt.Sprintf(`select * from todo.task_bound where t_id = %d;`, tid)).Scan(&bounds)
	db.SQL().Raw(fmt.Sprintf(`select * from todo.task_commit where uid = %d and t_id = %d;`, uid, tid)).Scan(&commits)
	info := map[uint]*api.TaskInfoCondition{}

	for _, c := range bounds {
		if info[c.TCID] == nil {
			info[c.TCID] = new(api.TaskInfoCondition)
		}
		info[c.TCID].WantParams = append(info[c.TCID].WantParams, c.Param)
	}

	for _, c := range commits {
		if info[c.TCID] == nil {
			info[c.TCID] = new(api.TaskInfoCondition)
		}
		info[c.TCID].GotParams = append(info[c.TCID].GotParams, c.Param)
	}

	conds := []api.TaskInfoCondition{}

	for t, c := range info {
		conds = append(conds, api.TaskInfoCondition{
			Type:       t,
			GotParams:  c.GotParams,
			WantParams: c.WantParams,
		})
	}

	return api.TaskInfoResponse{
		Name:      task.Name,
		Desc:      task.Desc,
		Departure: task.Departure,
		Arrival:   task.Arrival,
		Conds:     conds,
	}, nil
}

const SQL_DetailedTask = "SELECT t.name AS topic_name, ta.name AS task_name, ta.`desc` AS task_desc, ta.departure, ta.arrival, ta.id AS task_id, tb.tc_id" +
	"\nFROM subscribe s" +
	"\nJOIN topic t ON s.tt_id = t.id" +
	"\nJOIN task ta ON t.id = ta.tt_id" +
	"\nJOIN task_bound tb ON ta.id = tb.t_id" +
	"\nWHERE s.uid = %d;"

func (c *TaskController) GetTaskByUID(uid uint) (res []po.DetailedTask, err error) {
	temp := []po.DetailedTask{}
	err = db.SQL().RawExec(SQL_DetailedTask, uid).Scan(&temp).Error
	if err != nil {
		return
	}
	filter := map[uint]po.DetailedTask{}
	attr := map[uint][]uint{}
	for _, e := range temp {
		if filter[e.Id].Id == 0 {
			filter[e.Id] = e
			attr[e.Id] = append(attr[e.Id], e.CondType)
		} else {
			found := false
			for _, id := range attr[e.Id] {
				if id == e.CondType {
					found = true
					break
				}
			}
			if !found {
				attr[e.Id] = append(attr[e.Id], e.CondType)
			}
		}
	}
	for id, conds := range attr {
		cur := filter[id]
		res = append(res, po.DetailedTask{
			Id:        cur.Id,
			Name:      cur.Name,
			Desc:      cur.Desc,
			Departure: cur.Departure,
			Arrival:   cur.Arrival,
			Conds:     conds,
		})
	}
	return
}
