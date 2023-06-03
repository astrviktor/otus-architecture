package handlers

import (
	"auth/internal/storage"
	"auth/internal/token"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type ResponseToken struct {
	Token string `json:"token"`
}

type RequestToken struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) HandleCreateToken(ctx *fasthttp.RequestCtx) {
	if len(ctx.Request.Body()) == 0 {
		_, _ = ctx.WriteString("body is empty\n")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	body := ctx.Request.Body()
	requestToken := RequestToken{}

	if err := json.Unmarshal(body, &requestToken); err != nil {
		_, _ = ctx.WriteString("error in body\n")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ok, err := h.storage.CheckProfile(
		storage.Profile{
			Username: requestToken.Username,
			Password: requestToken.Password,
		})

	if err != nil {
		_, _ = ctx.WriteString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if !ok {
		_, _ = ctx.WriteString("username / password not found\n")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ok, err = h.storage.CheckGame(requestToken.Id, requestToken.Username)
	if err != nil {
		_, _ = ctx.WriteString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if !ok {
		_, _ = ctx.WriteString("game / username not found\n")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	token, err := token.GetToken(requestToken.Id, requestToken.Username)
	if err != nil {
		_, _ = ctx.WriteString(err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	WriteResponse(ctx, &ResponseToken{
		Token: token,
	})

	return
}
