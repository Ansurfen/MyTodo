package main

import (
	"MyTodo/conf"
	"MyTodo/middleware/driver/etcd/v1"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
	"github.com/tufanbarisyildirim/gonginx"
)

type LBOption struct {
	ETCD conf.ETCDOption `yaml:"etcd"`
}

func main() {
	parser, err := conf.NewNginxParser(conf.NginxConfOption{
		File: "nginx.conf",
	})
	if err != nil {
		panic(err)
	}

	opt := LBOption{}
	conf.
		New(conf.Option{File: "boot.yaml"}).
		MustRead().
		MustBind(&opt)

	etcdSrv := etcd.New(opt.ETCD)
	kv := etcdSrv.NewKV()

	for {
		res, err := kv.GetWithPrefix("/todo/service")
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		var copy []*gonginx.Upstream
		origin := parser.FindUpstreams()
		err = copier.Copy(&copy, &origin)
		if err != nil {
			panic(err)
		}

		copyUpstreamDict := make(map[string]*gonginx.Upstream)
		for _, up := range copy {
			copyUpstreamDict[up.UpstreamName] = up
			up.UpstreamServers = nil
		}

		for _, kvPair := range res.Kvs {
			metadata, _ := etcdSrv.ParseEndpoint(kvPair.Value)
			if upstream, ok := copyUpstreamDict[metadata.Group]; ok {
				upstream.AddServer(&gonginx.UpstreamServer{
					Address:    fmt.Sprintf("%s:%d", metadata.Host, metadata.Port),
					Parameters: make(map[string]string),
					Flags:      make([]string, 0),
				})
			}
		}

		if !reflect.DeepEqual(copy, origin) {
			log.Println("update")
			originUpstreamDict := make(map[string]*gonginx.Upstream)
			for _, up := range origin {
				originUpstreamDict[up.UpstreamName] = up
			}
			for k, v := range copyUpstreamDict {
				originUpstreamDict[k].UpstreamServers = v.UpstreamServers
			}
			parser.WriteAs("./temp.conf")
		}
		time.Sleep(5 * time.Second)
	}
}
