package proxy

import (
	fromserver "websocketexample/serverclient/from_server"

	"github.com/gorilla/websocket"
)

type GetTicker struct {
	server fromserver.FromServerInterface
}

func NewGetTickerModel(server fromserver.FromServerInterface) *GetTicker {
	return &GetTicker{
		server: server,
	}
}

func (g *GetTicker) GetTicker(clientws *websocket.Conn, last string, message string) error {
	_formserver := fromserver.NewGetTicker(last, message, clientws)
	_formserver.Run()
	return nil
}
