package handler

import (
	"log"
	"net/http"
	"websocketexample/serverclient/proxy"

	"github.com/gin-gonic/gin"
)

type GetTickerHandler struct {
	Proxy *proxy.GetTicker
}

func (g *GetTickerHandler) Handler(c *gin.Context) {
	upgrader := NewUpgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to upgrade to WebSocket " + err.Error()})
		return
	}

	request := &GetTickerRequest{}
	err = c.ShouldBindQuery(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request " + err.Error()})
		return
	}

	g.Proxy.GetTicker(conn, request.Id)

}
