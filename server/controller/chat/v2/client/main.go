package main

import (
	proto "MyTodo/api/v2/chat"
	"MyTodo/conf"
	"MyTodo/middleware/driver/etcd/v1"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const etcdDialPrefix = "etcd://192.168.127.130:2379/"

func main() {
	service := etcd.New(conf.ETCDOption{
		EndPoints:   []string{"192.168.127.130:2379"},
		DialTimeout: 5,
	})
	naming := service.NewNaming("/todo/service")
	resolver, err := naming.NewEtcdResolver()
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(etcdDialPrefix+naming.GetPathServerName("greeter"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(resolver),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &proto.HelloRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
