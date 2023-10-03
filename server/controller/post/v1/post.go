package controller

import (
	"MyTodo/dao"
	"MyTodo/engine/v1/db"
	"MyTodo/engine/v1/starter"
	"MyTodo/model/po/v1"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	ctx                 *gin.Context
	PostDao             dao.PostDao
	PostFavoriteDao     dao.PostFavoriteDao
	PostCommentDao      dao.PostCommentDao
	PostCommentV2Dao    dao.PostCommentV2
	PostCommentFavorite dao.PostCommentFavorite
}

func Get(ctx starter.TodoContext) *PostController {
	return &PostController{ctx: ctx.Context()}
}

func (c *PostController) CreatePost(post *po.Post) error {
	return c.PostDao.Create(post)
}

const SQL_DetailedPost = `
SELECT p.id, p.created_at, p.content, p.image, u.id AS uid, u.is_male, u.name AS username, COUNT(DISTINCT f.id) AS favoriteCnt, COUNT(DISTINCT c.id) AS commentCnt
FROM post p
LEFT JOIN user u ON p.uid = u.id
LEFT JOIN post_favorite f ON p.id = f.pid
LEFT JOIN post_comment c ON p.id = c.pid
WHERE p.deleted_at IS NULL
GROUP BY p.id
ORDER BY p.created_at DESC
LIMIT %d OFFSET %d; 
`

const SQL_DetailedPostV2 = `
SELECT p.id, p.created_at, p.content, p.image, u.id AS uid, u.is_male, u.name AS username, COUNT(DISTINCT f.id) AS favoriteCnt
FROM post p
LEFT JOIN user u ON p.uid = u.id
LEFT JOIN post_favorite f ON p.id = f.pid
WHERE p.deleted_at IS NULL
GROUP BY p.id
ORDER BY p.created_at DESC
LIMIT %d OFFSET %d; 
`

func (c *PostController) GetPostV2(page, count int) (ret []po.SnapshotPost, err error) {
	if page <= 0 {
		page = 1
	}
	if count < 0 {
		count = 0
	}
	err = db.SQL().Raw(fmt.Sprintf(SQL_DetailedPostV2, count, (page-1)*10)).Scan(&ret).Error
	if err != nil {
		return
	}
	for i := 0; i < len(ret); i++ {
		cnt, err := c.PostCommentV2Dao.CommentCount(int(ret[i].Id))
		if err == nil {
			ret[i].CommentCnt = int(cnt)
		}
	}
	return
}

func (c *PostController) GetPost(page, count int) (ret []po.SnapshotPost, err error) {
	if page <= 0 {
		page = 1
	}
	if count < 0 {
		count = 0
	}
	err = db.SQL().Raw(fmt.Sprintf(SQL_DetailedPost, count, (page-1)*10)).Scan(&ret).Error
	return
}

func (c *PostController) DelPost(pid int) error {
	return c.PostDao.DeleteByID(pid)
}

func (c *PostController) EditPost(p *po.Post) error {
	return c.PostDao.Edit(p)
}

func (c *PostController) CreatePostComment(pc *po.PostComment) error {
	return c.PostCommentDao.Create(pc)
}

func (c *PostController) GetPostComment(page, count int) (ret []po.PostComment, err error) {
	if page <= 0 {
		page = 1
	}
	if count < 0 {
		count = 0
	}
	return c.PostCommentDao.Get(page, count)
}

func (c *PostController) DelPostComment(pid, uid uint) error {
	return c.PostCommentDao.Del(int(pid), int(uid))
}

func (c *PostController) CreatePostFavorite(pid, uid uint) error {
	return c.PostFavoriteDao.Create(&po.PostFavorite{PID: pid, UID: uid})
}

func (c *PostController) DelPostFavorite(pid, uid int) error {
	return c.PostFavoriteDao.Del(pid, uid)
}

func (c *PostController) GetPostFavorite(pid int) (int64, error) {
	return c.PostFavoriteDao.Get(pid)
}

const SQLBatchCommentFavoriteCount = `SELECT cid, COUNT(*) AS count FROM post_comment_favorite WHERE cid IN (%s) GROUP BY cid;`

func (c *PostController) GetPostCommentFavoriteCount(cids []string) (res []po.DetailedPostCommentFavorite, err error) {
	cidf := []string{}
	for _, c := range cids {
		cidf = append(cidf, fmt.Sprintf(`'%s'`, c))
	}
	err = db.SQL().RawExec(SQLBatchCommentFavoriteCount, strings.Join(cidf, ",")).Scan(&res).Error
	return
}

const SQLGetPostDetail = `SELECT u.is_male AS isMale, u.name AS username, p.id, p.uid, p.content, p.image, COUNT(pf.id) AS favorite_count
FROM post p
JOIN user u ON p.uid = u.id
LEFT JOIN post_favorite pf ON p.id = pf.pid
WHERE p.id = %d
GROUP BY u.name, p.uid;`

func (c *PostController) GetPostDetail(id int) (res po.DetailedPost, err error) {
	err = db.SQL().RawExec(SQLGetPostDetail, id).Scan(&res).Error
	return
}
