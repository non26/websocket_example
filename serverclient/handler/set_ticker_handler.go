package handler

import (
	"log"
	"net/http"
	fromserver "websocketexample/serverclient/from_server"
	"websocketexample/serverclient/proxy"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type GetTickerHandler struct{}

func (g *GetTickerHandler) NewUpgrader() *websocket.Upgrader {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return &upgrader
}

func (g *GetTickerHandler) Handler(c *gin.Context) {
	upgrader := g.NewUpgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to upgrade to WebSocket " + err.Error()})
		return
	}
	defer conn.Close()

	request := &GetTickerRequest{}
	err = c.ShouldBindQuery(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request " + err.Error()})
		return
	}

	newformserver := fromserver.NewGetTicker(request.Last, request.Message, conn)
	newproxy := proxy.NewGetTickerModel(newformserver)
	newproxy.GetTicker(conn, request.Last, request.Message)
}
