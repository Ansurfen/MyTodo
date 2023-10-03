package main

import (
	"MyTodo/docs"
	"MyTodo/log"

	"MyTodo/engine/v1/cli"
	"MyTodo/engine/v1/db"
	"MyTodo/engine/v1/starter"
	middleware "MyTodo/middleware/handler"
	"MyTodo/routes"
	"MyTodo/utils/vfs"

	"go.uber.org/zap"
)

// @BasePath /
// @title MyTodo
// @version 1.0
// @description

func main() {
	log.Set(log.ZapAdaptor(zap.S().Desugar()))

	// creates cache directory for oss mode in local mode
	for _, dir := range []string{"assets/oss/user/profile"} {
		vfs.Mkdir(dir)
	}

	// enters oss mode
	vfs.Set(vfs.OSSAdapater(db.OSS()))

	for _, dir := range []string{"user", "chat", "post", "task"} {
		vfs.Mkdir(dir)
	}

	s := starter.New(starter.Option{
		Mode: cli.Option.Starter.Mode,
		SW: starter.OptionSwagger{
			Enabled:  cli.Option.Starter.SW.Enabled,
			Spec:     docs.SwaggerInfo,
			BasePath: cli.Option.Starter.SW.BasePath,
			Version:  cli.Option.Starter.SW.Version,
		},
		Middleware: starter.Middleware{middleware.CORS()},
		Registry: starter.RouteRegistry{
			"user":     routes.TodoRouter.InstallUserRoute,
			"chat":     routes.TodoRouter.InstallChatRoutes,
			"post":     routes.TodoRouter.InstallPostRoute,
			"topic":    routes.TodoRouter.InstallTopicRoute,
			"task":     routes.TodoRouter.InstallTaskRoute,
			"event":    routes.TodoRouter.InstallEventRoutes,
			"internal": routes.TodoRouter.InstallInternalRoute,
			"notify":   routes.TodoRouter.InstallNotificationRoute,
		},
	})
	// routes.TodoRouter.InstallCaptchaRoute(g)
	go s.WaitForShutdownSig()
	err := s.Bootstrap(cli.Option.Server)
	if err != nil {
		log.Fatal(err)
	}
}
