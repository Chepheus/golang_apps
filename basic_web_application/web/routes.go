package web

import (
	"net/http"

	"github.com/Chepheus/golang_apps/basic_web_application/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(rh handlers.RouteHandler) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(LogUrlToConsole)

	mux.Get("/", rh.HomeHandler)
	mux.Get("/about", rh.AboutHandler)

	return mux
}
