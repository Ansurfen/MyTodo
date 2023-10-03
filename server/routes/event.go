package routes

import (
	"MyTodo/engine/v1/starter"
	"MyTodo/service/event/v1"
)

type EventRoute struct{}

func (r *EventRoute) InstallEventRoutes(g *starter.MyTodoServerGroup) {
	g.GET("/event/qr", service.QrEvent)
}
