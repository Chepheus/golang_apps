package main

import (
	"fmt"
	"net/http"

	"github.com/Chepheus/golang_apps/basic_web_application/handlers"
	"github.com/Chepheus/golang_apps/basic_web_application/pkg"
	"github.com/Chepheus/golang_apps/basic_web_application/web"
)

const port = "8081"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello world!")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bites number:", n)
}

func main() {
	fmt.Println("Start...")

	appConfig := pkg.AppConfig{IsUseCache: true}
	templateRenderer := pkg.NewTemplateRenderer(appConfig)
	routeHandler := handlers.NewHandlers(appConfig, templateRenderer)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: web.Routes(routeHandler),
	}

	//http.HandleFunc("/hello", HelloWorld)

	fmt.Println("Listen port", port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
