package web

import (
	"net/http"

	"github.com/Chepheus/golang_apps/bookingns/pkg"
)

type Handler struct {
	templateRenderer pkg.TemplateRenderer
}

func NewHandler(templateRenderer pkg.TemplateRenderer) *Handler {
	return &Handler{templateRenderer: templateRenderer}
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.templateRenderer.RenderTemplate(w, "home.page.tmpl", nil)
}

func (h Handler) About(w http.ResponseWriter, r *http.Request) {
	h.templateRenderer.RenderTemplate(w, "about.page.tmpl", nil)
}

func (h Handler) Reservation(w http.ResponseWriter, r *http.Request) {
	h.templateRenderer.RenderTemplate(w, "reservation.page.tmpl", nil)
}
