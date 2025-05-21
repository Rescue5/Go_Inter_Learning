package auth

import (
	"HttpServer/extra"
	"HttpServer/pkg/req"
	"HttpServer/pkg/res"
	"fmt"
	"log"
	"net/http"
)

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
