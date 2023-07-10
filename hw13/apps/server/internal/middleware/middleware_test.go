package middleware

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"server/internal/handlers"
	"testing"
)

func TestCheckJWTEmptyBody(t *testing.T) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.SetRequestURI("/")

	handler := CheckJWT(func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("ok")
	})

	handler(&fasthttp.RequestCtx{
		Request:  *req,
		Response: *resp,
	})

	if resp.StatusCode() != fasthttp.StatusBadRequest {
		t.Errorf("CheckJWT returned incorrect status code: got %d, expected %d", resp.StatusCode(), fasthttp.StatusBadRequest)
	}

	body := resp.Body()
	expectedBody := "body is empty\n"
	if string(body) != expectedBody {
		t.Errorf("CheckJWT returned incorrect response body: got %s, expected %s", string(body), expectedBody)
	}
}

func TestCheckJWTInvalidBody(t *testing.T) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.SetRequestURI("/")
	req.SetBody([]byte("invalid json"))

	handler := CheckJWT(func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("ok")
	})

	handler(&fasthttp.RequestCtx{
		Request:  *req,
		Response: *resp,
	})

	if resp.StatusCode() != fasthttp.StatusBadRequest {
		t.Errorf("CheckJWT returned incorrect status code: got %d, expected %d", resp.StatusCode(), fasthttp.StatusBadRequest)
	}

	body := resp.Body()
	expectedBody := "error in body\n"
	if string(body) != expectedBody {
		t.Errorf("CheckJWT returned incorrect response body: got %s, expected %s", string(body), expectedBody)
	}
}

func TestCheckJWTUnauthorized(t *testing.T) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.SetRequestURI("/")
	req.SetBody([]byte(`{"foo": "bar"}`))

	handler := CheckJWT(func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("ok")
	})

	handler(&fasthttp.RequestCtx{
		Request:  *req,
		Response: *resp,
	})

	if resp.StatusCode() != fasthttp.StatusUnauthorized {
		t.Errorf("CheckJWT returned incorrect status code: got %d, expected %d", resp.StatusCode(), fasthttp.StatusUnauthorized)
	}

	body := resp.Body()
	expectedBody := "operation not unauthorized\n"
	if string(body) != expectedBody {
		t.Errorf("CheckJWT returned incorrect response body: got %s, expected %s", string(body), expectedBody)
	}
}

func TestCheckJWTAuthorized(t *testing.T) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	token := "valid_token"
	message := handlers.Message{Token: token}
	messageBytes, _ := json.Marshal(message)

	req.Header.SetMethod("POST")
	req.SetRequestURI("/")
	req.SetBody(messageBytes)

	handler := CheckJWT(func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("ok")
	})

	handler(&fasthttp.RequestCtx{
		Request:  *req,
		Response: *resp,
	})

	if resp.StatusCode() != fasthttp.StatusOK {
		t.Errorf("CheckJWT returned incorrect status code: got %d, expected %d", resp.StatusCode(), fasthttp.StatusOK)
	}

	body := resp.Body()
	expectedBody := "ok"
	if string(body) != expectedBody {
		t.Errorf("CheckJWT returned incorrect response body: got %s, expected %s", string(body), expectedBody)
	}
}
