package auth

import (
	"HttpServer/extra"
	"HttpServer/pkg/res"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Прочитать body
		var payload LoginRequest
		payload, err := GetLoginReq(writer, request, &payload)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(payload)
		// Сформировать ответ
		data := LoginResponse{Token: handler.Config.Auth.Secret}
		res.Json(writer, data, 200)
	}
}

func GetLoginReq(writer http.ResponseWriter, request *http.Request, payload *LoginRequest) (LoginRequest, error) {
	err := json.NewDecoder(request.Body).Decode(payload)

	if err != nil {
		res.Json(writer, "Password and Login Required", 402)
		return LoginRequest{}, extra.WrapError("Ошибка получения body в Login: ", err)
	}

	if payload.Password == "" || payload.Login == "" {
		err := errors.New(fmt.Sprintf("Login: %v, Password: %v", payload.Login, payload.Password))
		res.Json(writer, "Password and Login Required", 402)
		return LoginRequest{}, extra.WrapError("Не был получен логин или пароль", err)
	}

	return *payload, nil
}
