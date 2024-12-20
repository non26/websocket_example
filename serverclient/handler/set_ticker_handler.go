package handler

import (
	"log"
	"net/http"
	"websocketexample/serverclient/proxy"

	"github.com/gin-gonic/gin"
)

type SetTickerHandler struct {
	Proxy *proxy.SetTicker
}

func (s *SetTickerHandler) Handler(c *gin.Context) {
	upgrader := NewUpgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to upgrade to WebSocket " + err.Error()})
		return
	}
	// defer conn.Close() //this will close the connection

	request := &SetTickerRequest{}
	err = c.ShouldBindQuery(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request " + err.Error()})
		return
	}
	s.Proxy.SetTicker(conn, request.Id, request.Last, request.Message)
}
