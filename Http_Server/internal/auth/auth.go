package auth

import (
	"HttpServer/configs"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(r *http.ServeMux, deps AuthHandlerDeps) {
	handler := AuthHandler{deps.Config}
	r.HandleFunc("POST /auth/login", handler.Login())
	r.HandleFunc("POST /auth/register", handler.Register())
}
