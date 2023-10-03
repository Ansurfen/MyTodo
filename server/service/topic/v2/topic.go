package service

import (
	pb "MyTodo/api/v2/topic"
	"context"
	"fmt"
	"github.com/rookie-ninja/rk-grpc/v2/middleware/context"
	"google.golang.org/grpc"
)

type TopicService struct {
	pb.UnimplementedTopicServiceServer
}

var Handlers = pb.RegisterTopicServiceHandlerFromEndpoint

func Loader(server *grpc.Server) {
	pb.RegisterTopicServiceServer(server, &TopicService{})
}

func (server *TopicService) CreateTopic(ctx context.Context, _ *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {
	fmt.Println(rkgrpcctx.GetJwtToken(ctx))
	rkgrpcctx.GetLogger(ctx).Info("6666")
	return &pb.CreateTopicResponse{}, nil
}

func (server *TopicService) DeleteTopic(context.Context, *pb.DeleteTopicRequest) (*pb.DeleteTopicResponse, error) {
	return &pb.DeleteTopicResponse{}, nil
}
