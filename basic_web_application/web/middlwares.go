package web

import (
	"fmt"
	"net/http"
)

func LogUrlToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Current path:", request.URL.String())
		next.ServeHTTP(writer, request)
	})
}
