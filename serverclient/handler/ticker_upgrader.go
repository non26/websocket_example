package handler

import "github.com/gorilla/websocket"

func NewUpgrader() *websocket.Upgrader {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return &upgrader
}
