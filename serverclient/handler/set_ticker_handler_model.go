package handler

type GetTickerRequest struct {
	Last    string `form:"last"`
	Message string `form:"message"`
}
