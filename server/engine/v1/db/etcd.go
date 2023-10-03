package db

import (
	"MyTodo/engine/v1/cli"
	etcd "MyTodo/middleware/driver/etcd/v1"
)

var etcdClient *etcd.Client

func init() {
	etcdClient = etcd.New(cli.Option.ETCD)
}

func ETCD() *etcd.Client {
	return etcdClient
}
