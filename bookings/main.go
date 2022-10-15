package main

import (
	"fmt"
	"net/http"

	"github.com/Chepheus/golang_apps/bookingns/pkg"
	"github.com/Chepheus/golang_apps/bookingns/web"
)

const port = "8083"

func main() {
	appConfig := pkg.AppConfig{IsUseCache: false}
	tr := pkg.NewTemplateRenderer(appConfig)
	handler := web.NewHandler(tr)
	router := web.NewRouter(handler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}

	fmt.Println("Listen port", port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
