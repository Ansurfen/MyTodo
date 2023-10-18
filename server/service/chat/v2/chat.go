package service

import (
	pb "MyTodo/api/v2/chat"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func Loader(server *grpc.Server) {
	pb.RegisterChatServiceServer(server, &ChatService{})
}

var Handlers = pb.RegisterChatServiceHandlerFromEndpoint

type ChatService struct {
	pb.UnimplementedChatServiceServer
}

func (ChatService) DelChat(ctx context.Context, req *pb.DelChatRequest) (*pb.DelChatResponse, error) {
	fmt.Println(req.GetID())
	return &pb.DelChatResponse{}, nil
}

func (ChatService) EditChat(context.Context, *pb.EditChatRequest) (*pb.EditChatResponse, error) {
	return &pb.EditChatResponse{}, nil
}

func (ChatService) GetChat(context.Context, *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	return &pb.GetChatResponse{}, nil
}
