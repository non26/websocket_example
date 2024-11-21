package main

import (
	"websocketexample/serverclient/handler"
	"websocketexample/serverclient/hub"
	"websocketexample/serverclient/proxy"

	"github.com/gin-gonic/gin"
)

func registerRoute(route *gin.Engine, h *hub.ServerHub) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	_setProxy := proxy.SetTicker{Hub: h}
	_setTickerHandler := handler.SetTickerHandler{
		Proxy: &_setProxy,
	}
	_getProxy := proxy.GetTicker{Hub: h}
	_getTickerHandler := handler.GetTickerHandler{
		Proxy: &_getProxy,
	}
	route.GET("/set_ticker", _setTickerHandler.Handler)
	route.GET("/get_ticker", _getTickerHandler.Handler)
}

func main() {
	// TODO hub
	_hub := hub.NewServerHub()
	go _hub.Run()
	_route := gin.Default()
	registerRoute(_route, _hub)
	_route.Run(":8080")
}
