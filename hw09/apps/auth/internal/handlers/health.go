package handlers

import (
	"github.com/valyala/fasthttp"
)

type ResponseHealth struct {
	Status string `json:"status"`
}

func (h *Handler) HandleHealth(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	WriteResponse(ctx, &ResponseHealth{Status: "OK"})

	return
}
