package port

import "github.com/gorilla/websocket"

type FromServerInterface interface {
	Run() error
	GetId() string
	SetClientWs(clientws *websocket.Conn)
	// TODO add these methods:
	// ReadPump
	// WritePump
	// TickerPump
}
