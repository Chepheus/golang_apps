package handlers

import (
	"net/http"

	"github.com/Chepheus/golang_apps/basic_web_application/pkg"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pkg.RenderTemplate(w, "home.page.tmpl")
}
