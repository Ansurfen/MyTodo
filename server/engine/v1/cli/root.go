package cli

import (
	"MyTodo/conf"
	"MyTodo/engine/v1/starter"

	"github.com/spf13/cobra"
)

var (
	Option  TodoOption
	todoCmd = &cobra.Command{
		Use: "todo",
	}
)

type TodoOption struct {
	Server  conf.ServerOption `yaml:"server"`
	Mongo   conf.MongoOption  `yaml:"mongo"`
	SQL     conf.SQLOption    `yaml:"sql"`
	ES      conf.ESOption     `yaml:"es"`
	ETCD    conf.ETCDOption   `yaml:"etcd"`
	Redis   conf.RedisOption  `yaml:"redis"`
	Minio   conf.MinioOption  `yaml:"minio"`
	MQ      conf.MQOption     `yaml:"mq"`
	Starter starter.Option    `yaml:"starter"`
}

func init() {
	todoCmd.PersistentFlags().StringVarP(&Option.SQL.Username, "sql-username", "", "root", "")
	todoCmd.PersistentFlags().StringVarP(&Option.Server.Name, "name", "n", "todoService", "")
	todoCmd.PersistentFlags().IntVarP(&Option.Server.Port, "port", "p", 8080, "")
	todoCmd.PersistentFlags().StringVarP(&Option.Server.Host, "host", "", "localhost", "")
	todoCmd.PersistentFlags().StringSliceVarP(&Option.ES.Addresses, "es-address", "", []string{}, "")

	conf.
		New(conf.Option{
			File: "boot.yaml",
		}).
		MustRead().
		MustBind(&Option)
	err := todoCmd.Execute()
	if err != nil {
		panic(err)
	}
}
