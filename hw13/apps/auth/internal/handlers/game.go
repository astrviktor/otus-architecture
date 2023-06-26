package handlers

import (
	"auth/internal/storage"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type ResponseGameId struct {
	Id int `json:"id"`
}

func (h *Handler) HandleCreateGame(ctx *fasthttp.RequestCtx) {
	if len(ctx.Request.Body()) == 0 {
		_, _ = ctx.WriteString("body is empty\n")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	body := ctx.Request.Body()
	game := storage.Game{}

	if err := json.Unmarshal(body, &game); err != nil {
		_, _ = ctx.WriteString("error in body\n")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	id, err := h.storage.CreateGame(game)
	if err != nil {
		_, _ = ctx.WriteString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	WriteResponse(ctx, &ResponseGameId{
		Id: id,
	})

	return
}
