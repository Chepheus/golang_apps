package handlers

import (
	"net/http"
)

func (h RouteHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	h.templateRenderer.RenderTemplate(w, "home.page.tmpl", nil)
}
