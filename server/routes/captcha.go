package routes

import (
	"bytes"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

var ids *sync.Map

func init() {
	ids = &sync.Map{}
}

type CaptchaRoute struct{}

func (r *CaptchaRoute) InstallCaptchaRoute(group *gin.RouterGroup) {
	captchaRouter := group.Group("/captcha")
	{
		captchaRouter.GET("/new", captchaNew)
		captchaRouter.GET("/get/:source", captchaGet)
		captchaRouter.GET("/ls", captchaList)
		captchaRouter.POST("/verify", captchaVerify)
	}
}

func captchaNew(ctx *gin.Context) {
	id := captcha.New()
	ids.Store(id, true)
	ctx.String(http.StatusOK, id)
}

func captchaGet(ctx *gin.Context) {
	SetFile(ctx.Writer, ctx.Request)
}

func SetFile(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if setFile(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func setFile(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Captcha-Id", id)
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func captchaList(ctx *gin.Context) {
	keys := []string{}
	ids.Range(func(key, value any) bool {
		if v, ok := value.(bool); v && ok {
			keys = append(keys, key.(string))
		}
		return true
	})
	ctx.JSON(http.StatusOK, keys)
}

func captchaVerify(ctx *gin.Context) {
	id := ctx.PostForm("id")
	v := ctx.PostForm("v")
	if captcha.VerifyString(id, v) {
		ids.Store(id, false)
		ctx.String(http.StatusOK, "1")
		return
	}
	ctx.String(http.StatusOK, "0")
}
