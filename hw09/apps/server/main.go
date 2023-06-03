package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"server/internal/handlers"
	"server/internal/middleware"
)

const (
	Host = ""
	Port = 8008
)

func main() {
	handler := handlers.Handler{}

	r := router.New()
	r.GET("/health", handler.HandleHealth)
	r.POST("/message", middleware.CheckJWT(handler.HandleMessage))

	srv := &fasthttp.Server{
		Handler: r.Handler,
	}

	addr := fmt.Sprintf("%s:%d", Host, Port)

	err := srv.ListenAndServe(addr)
	fmt.Println(err)
}
