package hub

import "websocketexample/serverclient/port"

type ServerHub struct {
	tickerRegister   chan port.FromServerInterface
	tickerUnregister chan port.FromServerInterface

	tickerstore map[string]port.FromServerInterface
}

func NewServerHub() *ServerHub {
	return &ServerHub{
		tickerRegister:   make(chan port.FromServerInterface),
		tickerUnregister: make(chan port.FromServerInterface),
		tickerstore:      make(map[string]port.FromServerInterface),
	}
}

func (s *ServerHub) RegisterTicker(ticker port.FromServerInterface) {
	s.tickerRegister <- ticker
}

func (s *ServerHub) UnregisterTicker(ticker port.FromServerInterface) {
	s.tickerUnregister <- ticker
}

func (s *ServerHub) GetTicker(id string) (port.FromServerInterface, bool) {
	ticker, ok := s.tickerstore[id]
	return ticker, ok
}

func (s *ServerHub) Run() {
	for {
		select {
		case ticker := <-s.tickerRegister:
			s.tickerstore[ticker.GetId()] = ticker
			err := ticker.Run()
			if err != nil {
				println(err.Error())
			}
		case ticker := <-s.tickerUnregister:
			delete(s.tickerstore, ticker.GetId())
			// default:
			// 	println("no ticker")
		}
	}
}
