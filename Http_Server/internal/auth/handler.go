package auth

import (
	"HttpServer/configs"
	"HttpServer/extra"
	"HttpServer/pkg/req"
	"HttpServer/pkg/res"
	"fmt"
	"log"
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

func (handler *Handler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Прочитать body
		body, err := req.HandleBody[LoginRequest](writer, request)
		if err != nil {
			log.Println(extra.WrapError("Ошибка при чтении body:", err))
			return
		}
		fmt.Println(body)
		// Сформировать ответ
		data := LoginResponse{Token: handler.Config.Auth.Secret}
		res.Json(writer, data, 200)
	}
}

func (handler *Handler) Register() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[RegisterRequest](writer, request)
		if err != nil {
			log.Println(extra.WrapError("Ошибка чтения запроса регистрации", err))
			return
		}
		fmt.Println(body)

		data := RegisterResponse{Token: handler.Config.Auth.Secret}
		res.Json(writer, data, 200)
	}
}
