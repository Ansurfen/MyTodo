package service

import (
	"MyTodo/engine/v1/db"
	"MyTodo/engine/v1/starter"
	"MyTodo/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// @Summary QR Event
// @Description qr event
// @Tags Event
// @Accept json
// @Produce json
// @Success 101 {string} string "Switching Protocols"
// @Router /event/qr [get]
func QrEvent(ctx starter.TodoContext) {
	conn, err := upgrader.Upgrade(ctx.Context().Writer, ctx.Context().Request, nil)
	if err != nil {
		zap.S().Warn(err)
		return
	}
	messageType, message, err := conn.ReadMessage()
	if err != nil {
		zap.S().Warn(err)
		return
	}
	key := fmt.Sprintf("qr_%s", string(message))
	for {
		value := utils.RandString(8)
		db.Redis().Set(key, value, 10*time.Second)
		conn.WriteMessage(messageType, []byte(value))
		time.Sleep(10 * time.Second)
	}
}
