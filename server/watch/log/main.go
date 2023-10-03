package main

import (
	"MyTodo/conf"
	"MyTodo/utils"
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	cfg := conf.New(conf.Option{
		File: "./config.yaml",
	}).MustRead()
	ConfName := cfg.GetString("index", "log.index.json")
	if !utils.IsExist(ConfName) {
		os.WriteFile(ConfName, []byte("{}"), 0666)
	}
	go func() {
		c := cron.New()
		c.AddFunc(cfg.GetString("period"), func() {
			index := conf.New(conf.Option{
				File:      ConfName,
				Delimiter: "::",
			}).MustRead()
			logs := []string{}
			re := regexp.MustCompile(cfg.GetString("pattern"))
			paths := cfg.GetStringSlice("path")
			for _, root := range paths {
				filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.IsDir() {
						return nil
					}
					if re.MatchString(path) {
						log, err := filepath.Abs(filepath.Join(root, path))
						if err != nil {
							panic(err)
						}
						logs = append(logs, log)
					}
					return nil
				})
			}

			for _, filename := range logs {
				cnt := index.GetInt(filename)
				cur := 0
				file, err := os.Open(filename)
				if err != nil {
					panic(err)
				}
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					if cur >= cnt {
						fmt.Println(scanner.Text())
					}
					cur++
				}
				if cur != cnt {
					index.Set(filename, cur)
				}
			}
			err := index.WriteConfig()
			if err != nil {
				panic(err)
			}
		})
		c.Start()
	}()

	r := gin.Default()
	r.Run(":8085")
}
