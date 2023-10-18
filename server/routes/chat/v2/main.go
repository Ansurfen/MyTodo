package main

import (
	"MyTodo/conf"
	"MyTodo/engine/v2/starter"
	service "MyTodo/service/chat/v2"
	_ "embed"
)

//go:embed boot.yaml
var cfg []byte

type TopicOption struct {
	SQL conf.SQLOption `yaml:"sql"`
}

func main() {
	srv := starter.New(starter.Option{
		Config: cfg,
	})

	srv.ApplyIssuedMiddleware()

	opt := TopicOption{}

	if err := srv.DefaultUserConf().Unmarshal(&opt); err != nil {
		panic(err)
	}

	// db := sql.New(opt.SQL)
	// dao.SetDefault(db.DB)

	for name, info := range srv.GrpcOption.Service() {
		if info.Enabled {
			srv.NewThread(starter.LoadOption{
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
