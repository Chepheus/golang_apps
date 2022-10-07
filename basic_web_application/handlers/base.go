package handlers

import "github.com/Chepheus/golang_apps/basic_web_application/pkg"

type RouteHandler struct {
	container        *pkg.Container
	templateRenderer pkg.TemplateRenderer
}

func NewHandlers(container *pkg.Container, templateRenderer pkg.TemplateRenderer) RouteHandler {
	return RouteHandler{
		container:        container,
		templateRenderer: templateRenderer,
	}
}
