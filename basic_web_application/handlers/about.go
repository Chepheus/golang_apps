package handlers

import (
	"net/http"

	"github.com/Chepheus/golang_apps/basic_web_application/pkg"
)

func (h RouteHandler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	remoteIp := h.container.Session.GetString(r.Context(), "remote_ip")

	data := pkg.TemplateData{
		"title":    "About page",
		"remoteIp": remoteIp,
	}
	h.templateRenderer.RenderTemplate(w, "about.page.tmpl", &data)
}
