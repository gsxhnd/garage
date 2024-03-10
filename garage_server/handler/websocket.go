package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/gsxhnd/garage/garage_ffmpeg"
	"github.com/gsxhnd/garage/utils"
)

type WebsocketHandler interface {
	Ws(ctx *gin.Context)
}

type websocketHandler struct {
	logger   utils.Logger
	upgrader *websocket.Upgrader
}

func NewWebsocketHandler(l utils.Logger) WebsocketHandler {
	return &websocketHandler{
		logger: l,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *websocketHandler) Ws(ctx *gin.Context) {
	conn, err := h.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c, _ := garage_ffmpeg.NewVideoBatch(&garage_ffmpeg.VideoBatchOption{Exec: true})
	ob := c.GetExecBatch()
	go c.ExecuteBatch()

	go func() {
		fmt.Print("start listning...")
		for i := range ob.Observe() {
			fmt.Println("get ob item...")
			conn.WriteMessage(1, i.V.([]byte))
		}
	}()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Println("get exec batch")

		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
