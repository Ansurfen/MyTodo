package dao

import (
	"MyTodo/engine/v1/db"
	"MyTodo/model/po/v1"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	PostPageSize         = 10
	PostCommentPageSize  = 10
	PostFavoritePageSize = 10
)

type PostDao struct {
	SQLBaseDao[po.Post]
}

func (p *PostDao) Snapshot(page Pagination[po.SnapshotPost]) error {
	return page.SQLClause().Model(&po.Post{}).
		Select("posts.id, posts.created_at, posts.content, posts.image, users.id AS uid, users.is_male, users.name AS username, COUNT(DISTINCT post_favorites.id) AS favoriteCnt, COUNT(DISTINCT post_comments.id) AS commentCnt").
		Joins("LEFT JOIN users ON posts.uid = users.id").
		Joins("LEFT JOIN post_favorites ON posts.id = post_favorites.pid").
		Joins("LEFT JOIN post_comments ON posts.id = post_comments.pid").
		Where("posts.deleted_at IS NULL").
		Group("posts.id").
		Order("posts.created_at DESC").
		Find(&page.Data).Error
}

func (p *PostDao) Get(page, limit int) (ret []po.Post, err error) {
	err = db.SQL().Order("created_at DESC").Offset((page - 1) * PostPageSize).Limit(limit).Find(&ret).Error
	return
}

func (p *PostDao) GET(page Pagination[po.Post]) error {
	return page.SQLClause(PostPageSize).Order("created_at DESC").Find(&page.Data).Error
}

func (p *PostDao) Edit(c *po.Post) error {
	return db.SQL().Updates(c).Error
}

type PostCommentDao struct {
	po.PostComment
}

func (pc *PostCommentDao) Create(c *po.PostComment) error {
	return db.SQL().Create(c).Error
}

func (pc *PostCommentDao) Get(page, limit int) (ret []po.PostComment, err error) {
	err = db.SQL().Offset((page - 1) * PostCommentPageSize).Limit(limit).Find(&ret).Error
	return
}

func (pc *PostCommentDao) Edit(c *po.PostComment) error {
	return db.SQL().Updates(c).Error
}

func (pc *PostCommentDao) Delete(pid, uid int) error {
	return db.SQL().Where(PostCommentDao{
		PostComment: po.PostComment{
			PID: uint(pid),
			UID: uint(uid),
		},
	}).Update("deleted_at = ?", time.Now()).Error
}

type PostFavoriteDao struct {
	model po.PostFavorite
}

func (pf *PostFavoriteDao) Create(f *po.PostFavorite) error {
	return db.SQL().Create(f).Error
}

func (pf *PostFavoriteDao) Get(pid int) (ret int64, err error) {
	err = db.SQL().Model(pf.model).Where("pid = ?", pid).Count(&ret).Error
	return
}

func (pf *PostFavoriteDao) Delete(pid, uid int) error {
	return db.SQL().Unscoped().Delete(&po.PostFavorite{}, "pid = ? AND uid = ?", pid, uid).Error
}

type PostCommentV2 struct {
	MongoBaseDao[po.PostCommentV2]
}

func (c *PostCommentV2) Create(comment *po.PostCommentV2) error {
	res, err := c.Collection().InsertOne(context.Background(), comment)
	if err == nil {
		comment.ID = res.InsertedID.(primitive.ObjectID)
	}
	return err
}

func (c *PostCommentV2) CreateReply(id string, comment *po.PostCommentV2) error {
	oid, err := db.Mongo().BindID(id)
	if err != nil {
		return err
	}
	comment.ID = primitive.NewObjectID()
	_, err = c.Collection().UpdateOne(context.Background(), oid, bson.M{
		"$push": bson.M{"replies": comment},
	})
	return err
}

func (c *PostCommentV2) DeleteComment(uid int, id string) error {
	oid, err := db.Mongo().BindID(id)
	if err != nil {
		return err
	}
	oid["uid"] = uid
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}
	_, err = c.Collection().UpdateOne(context.Background(), oid, update)
	return err
}

func (c *PostCommentV2) DeleteCommentReply(uid int, commentID, replyID string) error {
	oid, err := db.Mongo().BindID(commentID)
	if err != nil {
		return err
	}
	replyOid, err := db.Mongo().BindID(replyID)
	if err != nil {
		return err
	}
	oid["replies._id"] = replyOid["_id"]
	oid["uid"] = uid
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}
	_, err = c.Collection().UpdateOne(context.Background(), oid, update)
	return err
}

func (c *PostCommentV2) EditComment(id string, comment *po.PostCommentV2) error {
	oid, err := db.Mongo().BindID(id)
	if err != nil {
		return err
	}
	oid["uid"] = comment.UID
	update := bson.M{
		"$set": bson.M{
			"content": comment.Content,
		},
	}
	_, err = c.Collection().UpdateOne(context.Background(), oid, update)
	return err
}

func (c *PostCommentV2) EditReply(commentID, replyID string, comment *po.PostCommentV2) error {
	oid, err := db.Mongo().BindID(commentID)
	if err != nil {
		return err
	}
	replyOid, err := db.Mongo().BindID(replyID)
	if err != nil {
		return err
	}
	oid["replies._id"] = replyOid["_id"]
	res, err := c.Collection().UpdateOne(context.Background(), oid, bson.M{
		"$set": bson.M{"replies.$.content": comment.Content},
	})
	fmt.Println(res)
	return err
}

func (c *PostCommentV2) Get(id, page, pageSize int) ([]po.PostCommentV2, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64((page - 1) * pageSize))

	cur, err := c.Collection().Find(context.Background(), bson.M{"pid": id, "deleted_at": time.Time{}}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var comments []po.PostCommentV2
	for cur.Next(context.Background()) {
		var comment po.PostCommentV2
		err = cur.Decode(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (c *PostCommentV2) CommentCount(id int) (int64, error) {
	return c.Collection().CountDocuments(context.TODO(), bson.M{
		"pid":        id,
		"deleted_at": time.Time{},
	})
}

type PostCommentFavorite struct{}

func (PostCommentFavorite) Create(f *po.PostCommentFavorite) error {
	return db.SQL().Create(f).Error
}

func (PostCommentFavorite) Delete(cid string, uid int) error {
	return db.SQL().Unscoped().Delete(&po.PostCommentFavorite{}, "cid = ? AND uid = ?", cid, uid).Error
}

func (PostCommentFavorite) Count(cid string) (res int64, err error) {
	err = db.SQL().Model(po.PostCommentFavorite{}).Where("cid = ?", cid).Count(&res).Error
	return
}

func (PostCommentFavorite) FindByUID(uid int) (res []po.PostCommentFavorite, err error) {
	err = db.SQL().Model(po.PostCommentFavorite{}).Where("uid = ?", uid).Find(&res).Error
	return
}
