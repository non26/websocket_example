package fromserver

import (
	"strconv"
	"time"
	"websocketexample/serverclient/hub"
	"websocketexample/serverclient/port"

	"github.com/gorilla/websocket"
)

type SetTicker struct {
	Id                string `json:"id"`
	Last              string `json:"last"`
	Message           string `json:"message"`
	Clientws          *websocket.Conn
	Hub               *hub.ServerHub
	is_clientws_close bool
}

func NewSetTicker(ticker *SetTicker) port.FromServerInterface {
	return &SetTicker{
		Id:       ticker.Id,
		Last:     ticker.Last,
		Message:  ticker.Message,
		Clientws: ticker.Clientws,
		Hub:      ticker.Hub,
	}
}

func (s *SetTicker) SetClientWs(clientws *websocket.Conn) {
	s.is_clientws_close = false
	s.Clientws = clientws
	go s.readPump()
}

func (s *SetTicker) GetLast() (int, error) {
	_last, err := strconv.Atoi(s.Last)
	if err != nil {
		return 0, err
	}
	return _last, nil
}

func (s *SetTicker) NewUpgrader() websocket.Upgrader {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return upgrader
}

func (s *SetTicker) readPump() {
	defer func() {
		s.Clientws.Close()
	}()
	for {
		_, _, err := s.Clientws.ReadMessage()
		if err != nil {
			println("read error-1", err.Error())
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure) {
				println("read error-2", err.Error())
				// break
			}
			break
		}
	}
}

func (s *SetTicker) GetId() string {
	return s.Id
}

func (s *SetTicker) Run() error {
	_last, err := s.GetLast()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	ticker_number := 0
	s.Clientws.SetPongHandler(func(appData string) error {
		return nil
	})
	// SetCloseHandler will be call if we read pump
	s.Clientws.SetCloseHandler(func(code int, text string) error {
		println("close handler")
		return nil
	})

	go s.readPump()

	for range ticker.C {
		ticker_number++
		if ticker_number > _last {
			s.Hub.UnregisterTicker(s)
			break
		}
		if s.Clientws == nil {
			println("clientws is nil")
			println(ticker_number)
			continue
		}
		println(ticker_number)
		if s.is_clientws_close {
			continue
		}
		err := s.Clientws.WriteMessage(websocket.TextMessage, []byte(s.Message+strconv.Itoa(ticker_number)))
		if err != nil {
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure) {
				s.Clientws.Close()
				println("close clientws")
				s.is_clientws_close = true
				continue
			}
			return err
		}
	}
	return nil
}
