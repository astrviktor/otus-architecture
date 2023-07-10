package main

import (
	"auth/internal/handlers"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	Host = ""
	Port = 8009
)

func main() {
	handler, _ := handlers.New()

	r := router.New()
	r.GET("/health", handler.HandleHealth)
	r.POST("/game", handler.HandleCreateGame)
	r.POST("/token", handler.HandleCreateToken)

	srv := &fasthttp.Server{
		Handler: r.Handler,
	}

	addr := fmt.Sprintf("%s:%d", Host, Port)

	err := srv.ListenAndServe(addr)
	fmt.Println(err)
}
