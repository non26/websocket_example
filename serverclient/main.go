package main

import (
	"websocketexample/serverclient/handler"

	"github.com/gin-gonic/gin"
)

func registerRoute(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	_handler := handler.GetTickerHandler{}
	route.GET("/get_ticker", _handler.Handler)
}

func main() {
	// TODO hub
	_route := gin.Default()
	registerRoute(_route)
	_route.Run(":8080")
}
