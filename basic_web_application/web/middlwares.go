package web

import (
	"fmt"
	"net/http"
)

func (w Web) LogUrlToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Current path:", request.URL.String())
		next.ServeHTTP(writer, request)
	})
}

func (w Web) SessionLoad(next http.Handler) http.Handler {
	return w.container.Session.LoadAndSave(next)
}
