package proxy

import (
	fromserver "websocketexample/serverclient/from_server"
	"websocketexample/serverclient/hub"

	"github.com/gorilla/websocket"
)

type SetTicker struct {
	Hub *hub.ServerHub
}

func NewSetTickerModel(hub *hub.ServerHub) *SetTicker {
	return &SetTicker{
		Hub: hub,
	}
}

func (s *SetTicker) SetTicker(clientws *websocket.Conn, id string, last string, message string) error {

	ticker := &fromserver.SetTicker{
		Id:       id,
		Last:     last,
		Message:  message,
		Clientws: clientws,
		Hub:      s.Hub,
	}
	_formserver := fromserver.NewSetTicker(ticker)
	s.Hub.RegisterTicker(_formserver)
	return nil
}
