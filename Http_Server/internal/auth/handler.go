package auth

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

func NewHandler(r *http.ServeMux, deps HandlerDeps) {
	handler := Handler{deps.Config}
	r.HandleFunc("POST /auth/login", handler.Login())
	r.HandleFunc("POST /auth/register", handler.Register())
}
