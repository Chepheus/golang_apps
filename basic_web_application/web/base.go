package web

import "github.com/Chepheus/golang_apps/basic_web_application/pkg"

type Web struct {
	container *pkg.Container
}

func NewWeb(c *pkg.Container) Web {
	return Web{container: c}
}
