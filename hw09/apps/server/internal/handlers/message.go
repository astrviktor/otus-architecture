package handlers

import "github.com/valyala/fasthttp"

type Message struct {
	Token       string        `json:"token"`
	GameID      int           `json:"gameId"`
	User        string        `json:"user"`
	ObjectID    int           `json:"objectId"`
	OperationID int           `json:"operationId"`
	Args        []interface{} `json:"args"`
}

func (h *Handler) HandleMessage(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString("operation authorized\n")
	ctx.SetStatusCode(fasthttp.StatusOK)

	return
}
