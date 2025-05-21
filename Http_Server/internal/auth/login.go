package auth

import (
	"HttpServer/extra"
	"HttpServer/pkg/req"
	"HttpServer/pkg/res"
	"fmt"
	"log"
	"net/http"
)

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
