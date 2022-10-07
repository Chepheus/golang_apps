package handlers

import (
	"net/http"
)

func (h RouteHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	h.container.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)

	h.templateRenderer.RenderTemplate(w, "home.page.tmpl", nil)
}
