package service

import (
	pb "MyTodo/api/v2/topic"
	userController "MyTodo/controller/user/v2"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

var Handlers = pb.RegisterTopicServiceHandlerFromEndpoint

func Loader(server *grpc.Server) {
	pb.RegisterTopicServiceServer(server, &TopicService{})
}

type TopicService struct {
	pb.UnimplementedTopicServiceServer
}

func (server *TopicService) CreateTopic(
	ctx context.Context,
	req *pb.CreateTopicRequest) (
	*pb.CreateTopicResponse, error) {
	uc, err := userController.Get(ctx)
	if err != nil {
		return &pb.CreateTopicResponse{}, err
	}
	fmt.Println(uc.User)
	// rkgrpcctx.GetLogger(ctx).Info("6666")
	return &pb.CreateTopicResponse{}, nil
}

func (server *TopicService) DeleteTopic(
	ctx context.Context,
	req *pb.DeleteTopicRequest) (
	*pb.DeleteTopicResponse, error) {
	return &pb.DeleteTopicResponse{}, nil
}
