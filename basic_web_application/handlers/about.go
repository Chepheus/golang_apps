package handlers

import (
	"net/http"
)

func (h RouteHandler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	h.templateRenderer.RenderTemplate(w, "about.page.tmpl")
}
