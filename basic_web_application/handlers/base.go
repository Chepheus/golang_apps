package handlers

import "github.com/Chepheus/golang_apps/basic_web_application/pkg"

type RouteHandler struct {
	appConfig        pkg.AppConfig
	templateRenderer pkg.TemplateRenderer
}

func NewHandlers(appConfig pkg.AppConfig, templateRenderer pkg.TemplateRenderer) RouteHandler {
	return RouteHandler{
		appConfig:        appConfig,
		templateRenderer: templateRenderer,
	}
}
