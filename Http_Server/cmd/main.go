package main

import (
	"HttpServer/configs"
	"HttpServer/db"
	"HttpServer/internal/auth"
	"HttpServer/internal/link"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(*conf)
	router := http.NewServeMux()

	// Repositories
	linkRepo := link.NewRepository(database)

	// Handlers
	auth.NewHandler(router, auth.HandlerDeps{Config: conf})
	link.NewHandler(router, link.HandlerDeps{
		Config:     conf,
		Repository: linkRepo,
	})

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
