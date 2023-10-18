package controller

import (
	"MyTodo/dao"
	"MyTodo/engine/v1/db"
	"MyTodo/engine/v1/starter"
	"MyTodo/model/bo/v1"
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

const SQLDetailedPostV2 = `SELECT p.id, p.created_at, p.content, p.image, u.id AS uid, u.is_male, u.name AS username, COUNT(DISTINCT f.id) AS favoriteCnt,
CASE WHEN EXISTS (
	SELECT 1 FROM post_favorite WHERE pid = p.id AND uid = %d AND deleted_at IS NULL
) THEN TRUE ELSE FALSE END AS is_favorite
FROM post p
LEFT JOIN user u ON p.uid = u.id
LEFT JOIN post_favorite f ON p.id = f.pid
WHERE p.deleted_at IS NULL
GROUP BY p.id
ORDER BY p.created_at DESC
LIMIT %d OFFSET %d;`

func (c *PostController) GetPostV2(uid, page, count int) (ret []bo.SnapshotPost, err error) {
	if page <= 0 {
		page = 1
	}
	if count < 0 {
		count = 0
	}

	err = db.SQL().Table("post").
		Select("post.id, post.created_at, post.content, post.image, user.id AS uid, user.is_male, user.name AS username, COUNT(DISTINCT post_favorite.id) AS favoriteCnt, CASE WHEN EXISTS (SELECT 1 FROM post_favorite WHERE pid = post.id AND uid = ? AND deleted_at IS NULL) THEN TRUE ELSE FALSE END AS is_favorite", uid).
		Joins("LEFT JOIN user ON post.uid = user.id").
		Joins("LEFT JOIN post_favorite ON post.id = post_favorite.pid").
		Where("post.deleted_at IS NULL").
		Group("post.id").
		Order("post.created_at DESC").
		Limit(count).
		Offset((page - 1) * 10).
		Find(&ret).Error
	// err = db.SQL().Raw(fmt.Sprintf(SQLDetailedPostV2, uid, count, (page-1)*10)).Scan(&ret).Error
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

func (c *PostController) GetPost(page, count int) (ret []bo.SnapshotPost, err error) {
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
	return c.PostCommentDao.Delete(int(pid), int(uid))
}

func (c *PostController) CreatePostFavorite(pid, uid uint) error {
	return c.PostFavoriteDao.Create(&po.PostFavorite{PID: pid, UID: uid})
}

func (c *PostController) DelPostFavorite(pid, uid int) error {
	return c.PostFavoriteDao.Delete(pid, uid)
}

func (c *PostController) GetPostFavorite(pid int) (int64, error) {
	return c.PostFavoriteDao.Get(pid)
}

const SQLBatchCommentFavoriteCount = `SELECT cid, COUNT(*) AS count FROM post_comment_favorite WHERE cid IN (%s) GROUP BY cid;`

func (c *PostController) GetPostCommentFavoriteCount(cids []string) (res []po.DetailedPostCommentFavorite, err error) {
	if len(cids) == 0 {
		return
	}
	cidf := []string{}
	for _, c := range cids {
		cidf = append(cidf, fmt.Sprintf(`'%s'`, c))
	}
	err = db.SQL().Table("post_comment_favorite").
		Select("cid, COUNT(*) AS count").
		Where("cid IN (?)", strings.Join(cidf, ",")).
		Group("cid").
		Find(&res).Error
	// err = db.SQL().RawExec(SQLBatchCommentFavoriteCount, strings.Join(cidf, ",")).Scan(&res).Error
	return
}

const SQLGetPostDetail = `SELECT u.is_male AS isMale, u.name AS username, p.id, p.uid, p.content, p.image, COUNT(pf.id) AS favorite_count,
CASE WHEN EXISTS (
	SELECT 1 FROM post_favorite WHERE pid = p.id AND uid = %d AND deleted_at IS NULL
) THEN TRUE ELSE FALSE END AS is_favorite
FROM post p
JOIN user u ON p.uid = u.id
LEFT JOIN post_favorite pf ON p.id = pf.pid
WHERE p.id = %d
GROUP BY u.name, p.uid;`

func (c *PostController) GetPostDetail(uid, id int) (res po.DetailedPost, err error) {
	// err = db.SQL().RawExec(SQLGetPostDetail, uid, id).Scan(&res).Error
	db.SQL().Select("user.is_male AS isMale, user.name AS username, post.id, post.uid, post.content, post.image, COUNT(post_favorite.id) AS favorite_count, CASE WHEN EXISTS (SELECT 1 FROM post_favorite WHERE pid = post.id AND uid = ? AND deleted_at IS NULL) THEN TRUE ELSE FALSE END AS is_favorite", uid).
		Joins("JOIN user ON post.uid = user.id").
		Joins("LEFT JOIN post_favorite ON post.id = post_favorite.pid").
		Where("post.id = ?", id).
		Group("user.name, post.uid").
		First(&res)
	return
}
