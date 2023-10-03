package fileController

import (
	"MyTodo/engine/v1/db"
	"context"
	"fmt"
	"image/color"

	"github.com/gin-gonic/gin"
	"github.com/issue9/identicon/v2"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

var (
	size = 128
	back = color.RGBA{R: 255, G: 0, B: 0, A: 100}
	fore = color.RGBA{R: 0, G: 255, B: 255, A: 100}
	ii1  = identicon.New(identicon.Style1, size, back, fore)
	ii2  = identicon.New(identicon.Style2, size, back, fore)
)

const PathAvatar = "/profile/%d.png"

type FileController struct {
	ctx *gin.Context
}

func (c *FileController) Profile(id int) {
	err := db.OSS().MakeBucket("user")
	if err != nil {
		zap.S().Error(err)
	}
	db.OSS().Put("user", fmt.Sprintf(PathAvatar, id), fmt.Sprintf("./assets/user/%d.png", id), minio.PutObjectOptions{
		ContentType: "application/png",
	})
	// fmt.Println(db.OSS().PresignedGetObject(context.TODO(), "user", fmt.Sprintf(PathAvatar, id), time.Minute*5, url.Values{}))
	db.OSS().Client.FGetObject(context.TODO(), "user", fmt.Sprintf(PathAvatar, id), "./xx.png", minio.GetObjectOptions{})
}
