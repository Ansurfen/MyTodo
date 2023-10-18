package fileController

import (
	"MyTodo/engine/v1/cli"
	"MyTodo/engine/v1/starter"
	"MyTodo/utils"
	"MyTodo/utils/vfs"
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/issue9/identicon/v2"
	"github.com/o1egl/govatar"
)

var (
	size = 128
	back = color.RGBA{R: 255, G: 0, B: 0, A: 100}
	fore = color.RGBA{R: 0, G: 255, B: 255, A: 100}
	ii1  = identicon.New(identicon.Style1, size, back, fore)
	ii2  = identicon.New(identicon.Style2, size, back, fore)
)

type FileController struct {
	ctx starter.TodoContext
}

const PathAvatar = "user:/profile/%d.png"

func Get(ctx starter.TodoContext) *FileController {
	return &FileController{ctx: ctx}
}

func (c *FileController) Profile(id int) {
	switch cli.Option.Starter.Profile {
	case "govatar":
		c.profileGovatar(id)
	case "ii1":
		c.profileIdenticonStyle1(id)
	case "ii2":
		c.profileIdenticonStyle2(id)
	default:
		c.profileGovatar(id)
	}
}

func (c *FileController) profileGovatar(id int) {
	dir := fmt.Sprintf(PathAvatar, id)
	if !vfs.IsExist(dir) {
		err := govatar.GenerateFileForUsername(govatar.MALE, fmt.Sprintf("%d", id), dir)
		if err != nil {
			c.ctx.Throw(err)
			return
		}
	}
	c.ctx.File(dir)
}

func (c *FileController) profileIdenticonStyle1(id int) {
	dir := fmt.Sprintf(PathAvatar, id)
	if !vfs.IsExist(dir) {
		img := ii1.Make([]byte(fmt.Sprintf("%d", id)))
		var err error
		if vfs.IsLoacl() {
			dir, err = vfs.Pathf(dir)
			if err != nil {
				c.ctx.Throw(err)
				return
			}
		} else {
			bucket, obejct, err := vfs.Objectf(dir)
			if err != nil {
				c.ctx.Throw(err)
				return
			}
			dir, err = vfs.Pathf(fmt.Sprintf("oss:%s/%s", bucket, obejct))
			if err != nil {
				c.ctx.Throw(err)
				return
			}
		}
		fi, err := os.Create(dir)
		if err != nil {
			c.ctx.Throw(err)
			return
		}
		err = png.Encode(fi, img)
		if err != nil {
			c.ctx.Throw(err)
			return
		}
		err = fi.Close()
		if err != nil {
			c.ctx.Throw(err)
			return
		}
		data, err := os.ReadFile(dir)
		if err != nil {
			c.ctx.Throw(err)
			return
		}
		err = vfs.WriteFile(fmt.Sprintf(PathAvatar, id), data)
		if err != nil {
			c.ctx.Throw(err)
			return
		}
	}
	c.ctx.File(fmt.Sprintf(PathAvatar, id))
}

func (c *FileController) profileIdenticonStyle2(id int) {
	dir := fmt.Sprintf(PathAvatar, id)
	if !utils.IsExist(dir) {
		img := ii2.Make([]byte(fmt.Sprintf("%d", id)))
		dir, err := vfs.Pathf(dir)
		if err != nil {
			c.ctx.Throw(err)
			return
		}
		fi, err := os.Create(dir)
		if err != nil {
			c.ctx.Throw(err)
			return
		}
		png.Encode(fi, img)
		fi.Close()
	}
	c.ctx.File(dir)
}
