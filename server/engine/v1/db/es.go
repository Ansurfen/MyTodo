package db

import (
	"MyTodo/engine/v1/cli"
	es "MyTodo/middleware/driver/es/v1"
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v7/esutil"
)

var esClient *es.TodoES

func init() {
	esClient = es.New(cli.Option.ES)
	logData := map[string]interface{}{
		"message": "This is a log message",
		"level":   "info",
	}

	// 使用Index API将日志内容传输到Elasticsearch
	res, err := esClient.Index(
		"user-log", // 替换为你的索引名称
		esutil.NewJSONReader(logData),
		esClient.Index.WithContext(context.Background()),
	)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	log.Printf("Indexed document ID: %s", res.String())
}

func ES() *es.TodoES {
	return esClient
}
