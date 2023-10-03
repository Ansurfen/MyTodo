package service

import (
	api "MyTodo/api/v1/user"
	fileController "MyTodo/controller/file/v1"
	userController "MyTodo/controller/user/v1"
	"MyTodo/engine/v1/starter"
	interfaces "MyTodo/interface"
	"MyTodo/utils"
	"fmt"
	"path/filepath"
	"strconv"

	"gorm.io/gorm"
)

// @Summary User profile
// @Description get user profile
// @Tags User
// @Param id path string true "id"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /user/profile/{id} [get]
func UserProfile(ctx starter.TodoContext) {
	fileController.
		Get(ctx).
		Profile(ctx.ParamInt("id"))
}

// @Summary User information
// @Description get user information
// @Tags User
// @Param id path string true "id"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /user/info/{id} [get]
func UserInfo(ctx starter.TodoContext, req api.UserInfoRequest) (interfaces.Response, error) {
	id := ctx.Context().Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	uc, err := userController.GetBlank(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	u, err := uc.UserDao.FindByID(i)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.UserInfoResponse{
		Name: u.Name,
	}, nil
}

// @Summary User sign
// @Description sign and login
// @Tags User
// @Param data body api.UserSignRequest true "UserSignRequest"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /user/sign [post]
func UserSign(ctx starter.TodoContext, req api.UserSignRequest) (interfaces.Response, error) {
	uc, err := userController.GetBlank(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	user, err := uc.FindByEmail(req.Email)
	if err != gorm.ErrRecordNotFound && err != nil {
		return ctx.ThrowWithResult(err)
	}
	if !user.Exist() {
		user = user.New(req.Email, req.Password)
		if err = uc.CreateUser(&user); err != nil {
			return ctx.ThrowWithResult(err)
		}
	} else if !user.Login(req.Password) {
		return ctx.ThrowWithResult(utils.ErrPassword)
	}
	jwt, err := utils.ReleaseToken(user.ID)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.UserSignResposne{
		JWT: jwt,
	}, nil
}

// @Summary Get user detailed information
// @Description Get user detailed information
// @Tags User
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /user/get [post]
func UserGet(ctx starter.TodoContext, req api.UserGetRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return ctx.ThrowWithResult(err)
	}
	return api.UserGetResponse{
		Id:        int(uc.User.ID),
		Name:      uc.User.Name,
		Email:     uc.User.Email,
		Telephone: uc.User.Telephone.String,
	}, nil
}

// @Summary Edit user information
// @Description Edit user information
// @Tags User
// @Accept multipart/form-data
// @Param profile formData file true "profile"
// @Param name formData string true "name"
// @Param email formData string true "email"
// @Param telephone formData string true "telephone"
// @Param x-token header string true "user jwt"
// @Success 200 {string} Success
// @failure 200 {object} string
// @Router /user/edit [post]
func UserEdit(ctx starter.TodoContext, req api.UserEditRequest) (interfaces.Response, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return api.UserEditResponse{}, utils.ErrUserNotFound
	}
	u := uc.User
	if req.Profile != nil {
		// fp, _ := req.Profile.Open()
		// db.OSS().Client.PutObject(context.Background(), "", "", fp, req.Profile.Size, minio.PutObjectOptions{
		// 	ContentType: req.Profile.Header.Get("Content-Type"),
		// })
		ctx.Context().SaveUploadedFile(req.Profile,
			filepath.Join("./assets/user/",
				fmt.Sprintf("%d.png", u.ID)))
	}
	// vaildates whether user's field is legal

	if len(req.Email) > 0 && req.Email != u.Email {
		u.Email = req.Email
	}

	if len(req.Telephone) > 0 && req.Telephone != u.Telephone.String {
		u.Telephone.String = req.Telephone
		u.Telephone.Valid = true
	}

	if len(req.Name) > 0 && req.Name != u.Name {
		u.Name = req.Name
	}

	err = uc.UserDao.UpdateUser(u.ID, u)
	if err != nil {
		return api.UserEditResponse{}, err
	}
	return api.UserEditResponse{}, nil
}
