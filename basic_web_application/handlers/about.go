package handlers

import (
	"net/http"

	"github.com/Chepheus/golang_apps/basic_web_application/pkg"
)

func (h RouteHandler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := pkg.TemplateData{
		"title": "About page",
	}
	h.templateRenderer.RenderTemplate(w, "about.page.tmpl", &data)
}
