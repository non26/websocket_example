package proxy

import (
	"errors"
	"websocketexample/serverclient/hub"

	"github.com/gorilla/websocket"
)

type GetTicker struct {
	Hub *hub.ServerHub
}

func NewGetTickerModel(hub *hub.ServerHub) *GetTicker {
	return &GetTicker{
		Hub: hub,
	}
}

func (s *GetTicker) GetTicker(clientws *websocket.Conn, id string) error {
	ticker, ok := s.Hub.GetTicker(id)
	if !ok {
		return errors.New("ticker not found")
	}
	ticker.SetClientWs(clientws)
	return nil
}
