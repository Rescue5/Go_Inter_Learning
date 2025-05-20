package main

import (
	"HttpServer/configs"
	"HttpServer/internal/auth"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Start listening...")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
