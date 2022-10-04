package main

import (
	"fmt"
	"net/http"

	"github.com/Chepheus/golang_apps/basic_web_application/handlers"
	"github.com/Chepheus/golang_apps/basic_web_application/pkg"
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
	handler := handlers.NewHandlers(appConfig, templateRenderer)

	http.HandleFunc("/hello", HelloWorld)

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/about", handler.AboutHandler)

	fmt.Println("Listen port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
