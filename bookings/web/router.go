package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{handler: handler}
}

func (r *Router) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", r.handler.Home)
	mux.Get("/about", r.handler.About)
	mux.Get("/reservation", r.handler.Reservation)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/ui/static/*", http.StripPrefix("/ui/static", fileServer))

	return mux
}
