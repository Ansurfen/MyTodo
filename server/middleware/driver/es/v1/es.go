package es

import (
	"MyTodo/conf"

	"github.com/elastic/go-elasticsearch/v7"
)

type TodoES struct {
	*elasticsearch.Client
}

func New(opt conf.ESOption) *TodoES {
	cli, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: opt.Addresses,
	})
	if err != nil {
		panic(err)
	}
	return &TodoES{cli}
}
