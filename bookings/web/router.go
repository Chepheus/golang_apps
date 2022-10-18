package web

import (
	"net/http"

	"github.com/Chepheus/golang_apps/bookingns/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	handler *handlers.Handler
}

func NewRouter(handler *handlers.Handler) *Router {
	return &Router{handler: handler}
}

func (r *Router) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", r.handler.Home)
	mux.Get("/about", r.handler.About)
	mux.Get("/reservation", r.handler.Reservation)
	mux.Post("/reservation", r.handler.CreateReservation)
	mux.Get("/json", r.handler.JsonResponse)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/ui/static/*", http.StripPrefix("/ui/static", fileServer))

	return mux
}
