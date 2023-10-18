package main

import (
	proto "MyTodo/api/v2/chat"
	"MyTodo/conf"
	"MyTodo/engine/v2/starter"
	"MyTodo/log"
	"MyTodo/middleware/driver/etcd/v1"

	"context"
	_ "embed"

	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkquery "github.com/rookie-ninja/rk-query"
	"google.golang.org/grpc"
)

//go:embed boot.yaml
var boot []byte

func regServer(server *grpc.Server) {
	proto.RegisterChatServiceServer(server, &Server{})
}

func main() {
	srv := starter.New(starter.Option{
		Config: boot,
	})

	etcdSrv := etcd.New(conf.ETCDOption{
		EndPoints: []string{"192.168.127.130:2379"},
	})

	namingSrv := etcdSrv.NewNaming("/todo/service")
	defer namingSrv.ReleaseAllEndpoint()

	for name, metadata := range srv.GrpcOption.Service() {
		if metadata.Enabled {
			srv.NewThread(starter.LoadOption{
				Name:         name,
				GrpcHandlers: starter.GrpcLoadFuncs{regServer},
				GWHandlers:   starter.GWLoadFuncs{proto.RegisterChatServiceHandlerFromEndpoint},
			})
			if err := namingSrv.RegisterEndpoint(etcd.Endpoint{
				Host:    rkentry.GlobalAppCtx.GetConfigEntry("my-config").GetString("host"),
				Name:    name,
				Port:    metadata.Port,
				Version: "1.0.0",
				Group:   "topic",
			}); err != nil {
				panic(err)
			}
		}
	}

	logger := rkentry.GlobalAppCtx.GetLoggerEntry("my-logger")
	logger.Info("This is my-logger")
	log.Set(log.ZapAdaptor(logger.Logger))

	// Try event
	eventEntry := rkentry.GlobalAppCtx.GetEventEntry("my-event")
	event := eventEntry.CreateEvent(rkquery.WithOperation("test"))
	event.AddPair("key", "value")
	event.Finish()

	srv.Bootstrap()

	srv.WaitForShutdownSig()
}

type Server struct {
	proto.UnimplementedChatServiceServer
}

func (s *Server) AddChat(ctx context.Context, req *proto.AddChatRequest) (*proto.AddChatResponse, error) {
	return &proto.AddChatResponse{Message: "topic"}, nil
}
