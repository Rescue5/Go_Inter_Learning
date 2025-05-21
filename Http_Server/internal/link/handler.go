package link

import (
	"HttpServer/configs"
	"net/http"
)

type Handler struct {
	*configs.Config
}

type HandlerDeps struct {
	*configs.Config
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := Handler{deps.Config}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("GET /{alias}", handler.Read())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}
