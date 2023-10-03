package service

import (
	api "MyTodo/api/v1/post"
	postController "MyTodo/controller/post/v1"
	userController "MyTodo/controller/user/v1"
	"MyTodo/engine/v1/starter"
	interfaces "MyTodo/interface"
	sql "MyTodo/middleware/driver/sql/v1"
	"MyTodo/model/bo/v1"
	"MyTodo/model/po/v1"
	"MyTodo/utils"
	"path/filepath"
	"strconv"
	"time"
)

// @Summary Create Post
// @Description create post
// @Tags Post
// @Accept multipart/form-data
// @Param content formData string true "content"
// @Param images formData file false "files"
// @Param x-token header string true "x-token"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/add [post]
func CreatePost(ctx starter.TodoContext, req api.CreatePostRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return api.CreateCommentPostResponse{}, utils.ErrUserNotFound
	}
	post := &po.Post{Content: req.Content, UID: uc.User.ID}
	for _, image := range req.Images {
		filename := utils.RandString(32) + filepath.Ext(image.Filename)
		absPath := filepath.Join(filename)
		post.Image = append(post.Image, po.Image{
			Name: image.Filename,
			Path: absPath,
		})
		ctx.Context().SaveUploadedFile(image, filepath.Join("./assets/post", filename))
	}
	pc := postController.Get(ctx)
	err = pc.CreatePost(post)
	if err != nil {
		return api.CreatePostResponse{}, err
	}
	return api.CreatePostResponse{}, nil
}

func GetPost(ctx starter.TodoContext, req api.GetPostRequest) (interfaces.Response, error) {
	posts, err := postController.Get(ctx).GetPost(req.Page, req.Count)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.GetPostResponse{
		Posts: posts,
	}, nil
}

// @Summary Post
// @Description post get
// @Tags Post
// @Param page query int true "page"
// @Param count query int true "count"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/get [get]
func GetPostV2(ctx starter.TodoContext, req api.GetPostRequest) (interfaces.Response, error) {
	posts, err := postController.Get(ctx).GetPostV2(req.Page, req.Count)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.GetPostResponse{
		Posts: posts,
	}, nil
}

// @Summary Get Post Detail
// @Description get post detail
// @Tags Post
// @Param id path int true "pid"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/detail/{id} [get]
func PostDetail(ctx starter.TodoContext, req interfaces.EmptyRequest) (interfaces.Response, error) {
	id, err := strconv.Atoi(ctx.Context().Param("id"))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	pc := postController.Get(ctx)
	res, err := pc.GetPostDetail(id)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.PostDetailResponse{
		DetailedPost: res,
	}, nil
}

func DeletePost(ctx starter.TodoContext, req api.DeletePostRequest) (interfaces.Response, error) {
	if err := postController.Get(ctx).DelPost(req.Pid); err != nil {
		return api.DeletePostResponse{}, err
	}
	return api.DeletePostResponse{}, nil
}

func EditPost(ctx starter.TodoContext, req api.EditPostRequest) (interfaces.Response, error) {
	if err := postController.Get(ctx).EditPost(&po.Post{
		Model: sql.Model{
			ID: req.PID,
		},
		UID:     req.UID,
		Content: req.Content,
		Image:   req.Images,
	}); err != nil {
		return api.EditPostResponse{}, err
	}
	return api.EditPostResponse{}, nil
}

func PostImage(ctx starter.TodoContext) {
	// controller.GetFileController(ctx)
	id := ctx.Context().Param("id")
	ctx.Context().File(filepath.Join("./assets/post", id))
}

// @Summary Favorite Post
// @Description favorite post
// @Tags Post
// @Accept multipart/form-data
// @Param pid formData int true "post id"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/favorite/add [post]
func FavoritePost(ctx starter.TodoContext, req api.FavoritePostRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	err = postController.Get(ctx).CreatePostFavorite(req.Pid, uc.User.ID)
	if err != nil {
		return api.FavoritePostResponse{}, err
	}
	return api.FavoritePostResponse{}, nil
}

// @Summary Unfavorite Post
// @Description unfavorite post
// @Tags Post
// @Accept multipart/form-data
// @Param pid formData int true "post id"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/favorite/del [post]
func UnfavoritePost(ctx starter.TodoContext, req api.UnfavoritePostRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	if err := postController.Get(ctx).DelPostFavorite(req.Pid, int(uc.User.ID)); err != nil {
		return api.UnfavoritePostResponse{}, err
	}
	return api.UnfavoritePostResponse{}, nil
}

// @Summary Favorite Post count
// @Description Get post favorite count
// @Tags Post
// @Accept multipart/form-data
// @Param pid formData int true "post id"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/favorite/get [post]
func GetPostFavorite(ctx starter.TodoContext, req api.GetPostFavoriteRequest) (interfaces.Response, error) {
	cnt, err := postController.Get(ctx).GetPostFavorite(req.Pid)
	if err != nil {
		return api.GetPostFavoriteResponse{}, err
	}
	return api.GetPostFavoriteResponse{Count: cnt}, nil
}

// @Summary Create Comment Post
// @Description Create Comment Post
// @Tags Post
// @Accept multipart/form-data
// @Param reply formData int true "reply"
// @Param pid formData int true "pid"
// @Param content formData string true "content"
// @Param files formData file false "images"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/add [post]
func CreateCommentPost(ctx starter.TodoContext, req api.CreateCommentPostRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	pc := postController.Get(ctx)
	err = pc.PostCommentV2Dao.Create(&po.PostCommentV2{
		CreatedAt: time.Now(),
		Reply:     req.Reply,
		PID:       uint(req.PID),
		UID:       uc.User.ID,
		Content:   []string{req.Content},
	})
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.CreateCommentPostResponse{}, nil
}

// @Summary Edit Comment Post
// @Description Edit Comment Post
// @Tags Post
// @Accept multipart/form-data
// @Param id formData string true "id"
// @Param content formData string true "content"
// @Param files formData file false "images"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/edit [post]
func EditCommentPost(ctx starter.TodoContext, req api.EditCommentPostRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return api.EditCommentPostResponse{}, nil
	}
	postController.Get(ctx).PostCommentV2Dao.EditComment(req.ID, &po.PostCommentV2{
		UID:     uc.User.ID,
		Content: []string{req.Content},
	})
	return api.EditCommentPostResponse{}, nil
}

// @Summary Delete Comment Post
// @Description Delete Comment Post
// @Tags Post
// @Accept multipart/form-data
// @Param id formData string true "id"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/del [post]
func DelCommentPost(ctx starter.TodoContext, req api.DelCommentPostRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return api.EditCommentPostResponse{}, nil
	}
	err = postController.Get(ctx).PostCommentV2Dao.DeleteComment(int(uc.User.ID), req.Id)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.DelCommentPostResponse{}, nil
}

// @Summary Get Comment Post
// @Description get Comment Post
// @Tags Post
// @Accept multipart/form-data
// @Param pid formData int true "pid"
// @Param page formData int true "page"
// @Param pageSize formData int true "pageSize"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/get [post]
func GetCommentPost(ctx starter.TodoContext, req api.GetCommentPostRequest) (interfaces.Response, error) {
	pc := postController.Get(ctx)
	res, err := pc.PostCommentV2Dao.Get(req.PID, req.Page, req.PageSize)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	var (
		uids      []uint
		replys    []uint
		cids      []string
		comments  []po.DetailedPostComment
		dict      = map[string]int64{}
		usersDict = map[uint]bo.User{}
		replyDict = map[uint]bo.User{}
	)
	for _, n := range res {
		uids = append(uids, n.UID)
		replys = append(replys, n.Reply)
	}
	uc, err := userController.GetBlank(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	users, err := uc.BatchFindUserByUID(uids)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	for _, u := range users {
		usersDict[u.ID] = u
	}
	replyers, err := uc.BatchFindUserByUID(replys)
	for _, u := range replyers {
		replyDict[u.ID] = u
	}
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	for _, n := range res {
		cids = append(cids, n.ID.Hex())
	}
	favorites, err := pc.GetPostCommentFavoriteCount(cids)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}

	for _, f := range favorites {
		dict[f.CID] = f.Count
	}
	for _, n := range res {
		cnt := int64(0)
		if v, ok := dict[n.ID.Hex()]; ok {
			cnt = v
		}
		username := ""
		if v, ok := usersDict[n.UID]; ok {
			username = v.Name
		}
		replyname := ""
		if v, ok := replyDict[n.Reply]; ok {
			replyname = v.Name
		}
		comments = append(comments, po.DetailedPostComment{
			PostCommentV2: n,
			Username:      username,
			ReplyName:     replyname,
			Favorite:      cnt,
		})
	}
	return api.GetCommentPostResponse{
		Comments: comments,
	}, nil
}

// @Summary Create Comment Reply
// @Description create comment reply
// @Tags Post
// @Accept multipart/form-data
// @Param id formData string true "id"
// @Param reply formData int true "reply id (if not user to reply, it's 0 in default)"
// @Param content formData string true "content"
// @Param files formData file false "images"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/reply/add [post]
func PostCommentReplyCreate(
	ctx starter.TodoContext,
	req api.PostCommentReplyCreateRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	pc := postController.Get(ctx)
	err = pc.PostCommentV2Dao.CreateReply(req.Id, &po.PostCommentV2{
		Content: []string{req.Content},
		UID:     uc.User.ID,
		Reply:   req.Reply,
	})
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.PostCommentReplyCreateResponse{}, nil
}

// @Summary Edit Comment Reply
// @Description edit comment reply
// @Tags Post
// @Accept multipart/form-data
// @Param comment_id formData string true "comment_id"
// @Param reply_id formData string true "reply_id"
// @Param content formData string true "content"
// @Param files formData file false "images"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/reply/edit [post]
func PostCommentReplyEdit(
	ctx starter.TodoContext,
	req api.PostCommentReplyEditRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	pc := postController.Get(ctx)
	pc.PostCommentV2Dao.EditReply(req.CommentID, req.ReplyID, &po.PostCommentV2{
		UID:     uc.User.ID,
		Content: []string{req.Content},
	})
	return api.PostCommentReplyEditResponse{}, nil
}

// @Summary Edit Comment Reply
// @Description edit comment reply
// @Tags Post
// @Accept multipart/form-data
// @Param comment_id formData string true "comment_id"
// @Param reply_id formData string true "reply_id"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/reply/del [post]
func PostCommentReplyDelete(
	ctx starter.TodoContext,
	req api.PostCommentReplyDeleteRequest) (
	interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	pc := postController.Get(ctx)
	pc.PostCommentV2Dao.EditReply(req.CommentID, req.ReplyID, &po.PostCommentV2{
		UID: uc.User.ID,
	})
	return api.PostCommentReplyDeleteResponse{}, nil
}

// @Summary Favorite comment
// @Description favorite comment
// @Tags Post
// @Accept multipart/form-data
// @Param comment_id formData string true "comment_id"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/favorite [post]
func PostCommentFavorite(ctx starter.TodoContext, req api.PostCommentFavoriteRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	pc := postController.Get(ctx)
	pc.PostCommentFavorite.Create(&po.PostCommentFavorite{
		CID: req.CommentID,
		UID: uc.User.ID,
	})
	return interfaces.BaseResponse{}, nil
}

// @Summary Unfavorite comment
// @Description unfavorite comment
// @Tags Post
// @Accept multipart/form-data
// @Param comment_id formData string true "comment_id"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/unfavorite [post]
func PostCommentUnfavorite(ctx starter.TodoContext, req api.PostCommentUnfavoriteRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	pc := postController.Get(ctx)
	err = pc.PostCommentFavorite.Delete(req.CommentID, int(uc.User.ID))
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return interfaces.BaseResponse{}, nil
}

// @Summary Get count of the comment favorite
// @Description get count of the comment favorite
// @Tags Post
// @Accept multipart/form-data
// @Param comment_id formData string true "comment_id"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /post/comment/favoriteCount [post]
func PostCommentFavoriteCount(ctx starter.TodoContext, req api.PostCommentFavoriteCountRequest) (interfaces.Response, error) {
	pc := postController.Get(ctx)
	cnt, err := pc.PostCommentFavorite.Count(req.CommentID)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.PostCommentFavoriteCountResponse{Count: cnt}, nil
}
