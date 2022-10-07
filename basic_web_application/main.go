package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Chepheus/golang_apps/basic_web_application/handlers"
	"github.com/Chepheus/golang_apps/basic_web_application/pkg"
	"github.com/Chepheus/golang_apps/basic_web_application/web"
	"github.com/alexedwards/scs/v2"
)

const port = "8081"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello world!")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bites number:", n)
}

func main() {
	fmt.Println("Start...")

	appConfig := setupAppConfig()
	session := setupSession()
	container := pkg.Container{AppConfig: appConfig, Session: session}

	templateRenderer := pkg.NewTemplateRenderer(container.AppConfig)
	routeHandler := handlers.NewHandlers(&container, templateRenderer)
	w := web.NewWeb(&container)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: w.Routes(routeHandler),
	}

	//http.HandleFunc("/hello", HelloWorld)

	fmt.Println("Listen port", port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func setupAppConfig() pkg.AppConfig {
	return pkg.AppConfig{IsUseCache: true}
}

func setupSession() *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	return session
}
