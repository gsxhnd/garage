package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gsxhnd/garage/utils"
)

type WebsocketHandler interface {
	Ws(c *websocket.Conn)
}

type websocketHandler struct {
	logger utils.Logger
	// upgrader *websocket.Upgrader
}

func NewWebsocketHandler(l utils.Logger) WebsocketHandler {
	return &websocketHandler{
		logger: l,
		// upgrader: &websocket.Upgrader{
		// 	ReadBufferSize:  1024,
		// 	WriteBufferSize: 1024,
		// },
	}
}

func (h *websocketHandler) Ws(c *websocket.Conn) {
	fmt.Println(c.Locals("Host"))
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)
		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
