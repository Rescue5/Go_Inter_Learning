package auth

import (
	"fmt"
	"net/http"
)

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Register")
	}
}
