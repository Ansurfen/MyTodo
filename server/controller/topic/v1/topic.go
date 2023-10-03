package topicController

import (
	"MyTodo/dao"
	"MyTodo/engine/v1/db"
	"MyTodo/engine/v1/starter"
	interfaces "MyTodo/interface"
	"MyTodo/model/po/v1"
	"MyTodo/utils"

	"github.com/gin-gonic/gin"
)

type TopicController struct {
	ctx          *gin.Context
	TopicDao     dao.TopicDao
	SubscribeDao dao.Subscribe
}

func Get(ctx starter.TodoContext) *TopicController {
	if con := ctx.Get(interfaces.ControllerTypeTopic); con != nil {
		if tc, ok := con.(*TopicController); ok {
			return tc
		}
	}
	return &TopicController{ctx: ctx.Context()}
}

func (c *TopicController) IsAdmin(ttid, uid int) bool {
	sub, err := c.SubscribeDao.FindOne(ttid, uid)
	if err != nil {
		return false
	}
	return sub.Perm == po.PermAdmin
}

func (c *TopicController) IsManager(ttid, uid int) bool {
	sub, err := c.SubscribeDao.FindOne(ttid, uid)
	if err != nil {
		return false
	}
	return sub.Perm == po.PermAdmin || sub.Perm == po.PermManager
}

func (c *TopicController) IsMemberOrManager(ttid, uid int) bool {
	sub, err := c.SubscribeDao.FindOne(ttid, uid)
	if err != nil {
		return false
	}
	return sub.Perm == po.PermMember || sub.Perm == po.PermManager
}

func (c *TopicController) IsJoin(ttid, uid int) bool {
	sub, err := c.SubscribeDao.FindOne(ttid, uid)
	if err != nil {
		return false
	}
	return sub.ID != 0
}

func (c *TopicController) CreateTopic(creator uint, name, desc string) (string, error) {
	inviteCode := utils.RandString(8)
	topic := &po.Topic{
		Creator:    creator,
		Name:       name,
		Desc:       desc,
		InviteCode: inviteCode,
	}
	if err := c.TopicDao.Create(topic); err != nil {
		return "", err
	}
	if err := c.SubscribeDao.Create(&po.Subscribe{
		TTID: topic.ID,
		UID:  creator,
		Perm: po.PermAdmin,
	}); err != nil {
		return "", err
	}
	return inviteCode, nil
}

const SQL_GetSubscribedTopic = `
SELECT *
FROM topic
WHERE id in (
  SELECT tt_id
  FROM subscribe
  WHERE uid = %d
  AND deleted_at IS NULL
)
AND deleted_at IS NULL;
`

func (c *TopicController) GetSubscribedTopic(uid int) (res []po.SubscribedTopic, err error) {
	err = db.SQL().RawExec(SQL_GetSubscribedTopic, uid).Scan(&res).Error
	return
}

func (c *TopicController) SubscribeTopic(uid uint, code string) error {
	topic, err := c.TopicDao.FindByInviteCode(code)
	if err != nil {
		return err
	}
	err = c.SubscribeDao.Create(&po.Subscribe{
		UID:  uid,
		TTID: topic.ID,
		Perm: po.PermMember,
	})
	if err != nil {
		return err
	}
	return nil
}

const SQL_GetSubscribedMemeber = "SELECT u.id, u.name FROM subscribe s LEFT JOIN `user` u ON u.id = s.uid  WHERE s.tt_id = %d;"

func (c *TopicController) GetSubscribedMemeber(ttid uint) (res []po.SubscribedMemeber, err error) {
	err = db.SQL().RawExec(SQL_GetSubscribedMemeber, ttid).Scan(&res).Error
	return
}
