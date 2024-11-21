package hub

import fromserver "websocketexample/serverclient/from_server"

type ServerHub struct {
	tickerRegister chan fromserver.FromServerInterface
}

// func NewServerHub() *ServerHub {
// 	return &ServerHub{
// 		tickerRegister: make(chan fromserver.FromServerInterface),
// 	}
// }

func (s *ServerHub) Run() {
	for {
		select {
		case ticker := <-s.tickerRegister:
			ticker.Run()
		}
	}
}
