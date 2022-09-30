package handlers

import (
	"net/http"

	"github.com/Chepheus/golang_apps/basic_web_application/helpers"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "about.page.tmpl")
}
