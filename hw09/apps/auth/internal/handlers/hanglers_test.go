package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	h := Handler{}

	ctx := fasthttp.RequestCtx{}
	h.HandleHealth(&ctx)

	assert.Equal(t, fasthttp.StatusOK, ctx.Response.StatusCode())
	assert.Equal(t, "{\"status\":\"OK\"}\n", string(ctx.Response.Body()))
}
