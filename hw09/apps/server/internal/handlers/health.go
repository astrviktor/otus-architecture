package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

type ResponseHealth struct {
	Status string `json:"status"`
}

func WriteResponse(ctx *fasthttp.RequestCtx, resp interface{}) {
	respBuf, err := json.Marshal(resp)
	if err != nil {
		log.Println(fmt.Sprintf("response marshal error: %s", err))
	}

	respBuf = append(respBuf, []byte("\n")...)
	ctx.Response.SetBody(respBuf)

	ctx.SetContentType("application/json; charset=utf-8")
}

func (h *Handler) HandleHealth(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	WriteResponse(ctx, &ResponseHealth{Status: "OK"})

	return
}
