package req

import (
	"HttpServer/extra"
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T

	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, extra.WrapError("Ошибка в декодировании тела запроса", err)
	}
	return payload, nil
}
