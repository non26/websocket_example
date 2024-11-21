package fromserver

import (
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type GetTicker struct {
	Last     string `json:"last"`
	Message  string `json:"message"`
	clientws *websocket.Conn
}

func (g *GetTicker) GetLast() (int, error) {
	_last, err := strconv.Atoi(g.Last)
	if err != nil {
		return 0, err
	}
	return _last, nil
}

func NewGetTicker(last string, message string, clientws *websocket.Conn) FromServerInterface {
	return &GetTicker{
		Last:     last,
		Message:  message,
		clientws: clientws,
	}
}

func (g *GetTicker) NewUpgrader() websocket.Upgrader {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return upgrader
}

func (g *GetTicker) readPump() {
	defer func() {
		g.clientws.Close()
	}()
	for {
		_, _, err := g.clientws.ReadMessage()
		if err != nil {
			println("read error-1")
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				println("read error-2", err.Error())
			}
			break
		}
	}
}

func (g *GetTicker) Run() error {
	_last, err := g.GetLast()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	ticker_number := 0
	g.clientws.SetPongHandler(func(appData string) error {
		return nil
	})
	// SetCloseHandler will be call if we read pump
	g.clientws.SetCloseHandler(func(code int, text string) error {
		println("close handler")
		return nil
	})

	go g.readPump()

	for range ticker.C {
		ticker_number++
		if ticker_number > _last {
			break
		}
		if g.clientws == nil {
			println("clientws is nil")
			println(ticker_number)
			continue
		}
		println(ticker_number)
		err := g.clientws.WriteMessage(websocket.TextMessage, []byte(g.Message))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err) {
				g.clientws.Close()
				println("close clientws")
				continue
			}
			return err
		}
	}
	return nil
}
