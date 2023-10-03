package routes

type TodoRoutes struct {
	CaptchaRoute
	TopicRoute
	UserRoute
	PostRoute
	TaskRoute
	EventRoute
	InternalRoute
	ChatRoute
	NotificationRoute
}

var TodoRouter = new(TodoRoutes)
