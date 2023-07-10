package middleware

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"server/internal/handlers"
	"server/internal/token"
)

func CheckJWT(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if len(ctx.Request.Body()) == 0 {
			_, _ = ctx.WriteString("body is empty\n")
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		body := ctx.Request.Body()
		message := handlers.Message{}

		if err := json.Unmarshal(body, &message); err != nil {
			_, _ = ctx.WriteString("error in body\n")
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		ok, err := token.CheckToken(message)
		if err != nil {
			_, _ = ctx.WriteString("error check token\n")
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		if !ok {
			_, _ = ctx.WriteString("operation not unauthorized\n")
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		h(ctx)
	}
}
