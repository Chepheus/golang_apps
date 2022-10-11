package main

import (
	"fmt"
	"net/http"

	"github.com/Chepheus/golang_apps/bookingns/pkg"
	"github.com/go-chi/chi/v5"
)

const port = "8083"

func main() {
	appConfig := pkg.AppConfig{IsUseCache: true}
	tr := pkg.NewTemplateRenderer(appConfig)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: Routes(tr),
	}

	fmt.Println("Listen port", port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Routes(templateRenderer pkg.TemplateRenderer) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templateRenderer.RenderTemplate(w, "home.page.tmpl", nil)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/ui/static/*", http.StripPrefix("/ui/static", fileServer))

	return mux
}
