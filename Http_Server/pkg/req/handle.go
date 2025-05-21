package req

import (
	"HttpServer/pkg/res"
	"net/http"
)

func HandleBody[T any](writer http.ResponseWriter, request *http.Request) (*T, error) {
	body, err := Decode[T](request.Body)
	if err != nil {
		res.Json(writer, err.Error(), 402)
		return nil, err
	}

	err = IsValid[T](body)
	if err != nil {
		res.Json(writer, err.Error(), 402)
		return nil, err
	}

	return &body, nil
}
