package main

import (
	"MyTodo/engine/v2/starter"
	service "MyTodo/service/topic/v2"
	_ "embed"
)

//go:embed boot.yaml
var cfg []byte

func main() {

	srv := starter.New(starter.Option{
		Config: cfg,
	})

	for name, info := range srv.GrpcOption.Service() {
		if info.Enabled {
			srv.Load(starter.LoadOption{
				Name:         name,
				GrpcHandlers: starter.GrpcLoadFuncs{service.Loader},
				GWHandlers:   starter.GWLoadFuncs{service.Handlers},
			})
			// if err := namingSrv.RegisterEndpoint(etcd.Endpoint{
			// 	Host:    rkentry.GlobalAppCtx.GetConfigEntry("my-config").GetString("host"),
			// 	Name:    name,
			// 	Port:    metadata.Port,
			// 	Version: "1.0.0",
			// 	Group:   "topic",
			// }); err != nil {
			// 	panic(err)
			// }
		}
	}

	srv.Bootstrap()

	srv.WaitForShutdownSig()
}
