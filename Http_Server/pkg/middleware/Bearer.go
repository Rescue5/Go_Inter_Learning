package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		bearerToken := strings.TrimPrefix(authorization, "Bearer ")
		fmt.Println(bearerToken)
		next.ServeHTTP(w, r)
	})
}
