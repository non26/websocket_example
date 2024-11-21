package handler

import (
	fromserver "websocketexample/serverclient/from_server"
)

type SetTickerRequest struct {
	Id      string `form:"id"`
	Last    string `form:"last"`
	Message string `form:"message"`
}

func (s *SetTickerRequest) ToSetTicker() *fromserver.SetTicker {
	return &fromserver.SetTicker{
		Id:      s.Id,
		Last:    s.Last,
		Message: s.Message,
	}
}

type GetTickerRequest struct {
	Id string `form:"id"`
}
