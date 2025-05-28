package main

import (
	"HttpServer/configs"
	"HttpServer/db"
	"HttpServer/internal/auth"
	"HttpServer/internal/link"
	"HttpServer/pkg/middleware"
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

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}
	fmt.Println("Start listening...")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
