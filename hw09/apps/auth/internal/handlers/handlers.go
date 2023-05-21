package handlers

import (
	"auth/internal/storage"
	"auth/internal/storage/memory"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

type Handler struct {
	storage storage.Storage
}

func New() (*Handler, error) {

	storage := memory.New()

	err := storage.Connect()
	if err != nil {
		return nil, err
	}

	return &Handler{
		storage: storage,
	}, nil
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
