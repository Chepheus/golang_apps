package main

import (
	"fmt"
	"golang_lessons/domain/repository"
	"golang_lessons/handler"
	"golang_lessons/logger"
	"golang_lessons/service"
	"net/http"

	"github.com/gorilla/mux"
)

func StartAuth(port string) {
	router := mux.NewRouter()

	db := getDbClient()
	userHandler := handler.UserHandler{Service: service.NewUserService(
		repository.NewUserRepository(db),
		repository.NewLoginRepository(db),
	)}

	router.HandleFunc("/auth/login", userHandler.Login).Methods(http.MethodGet)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", port), router)
	if err != nil {
		logger.Error(fmt.Sprintf("Server cant listen localhost:%s", port))
	}
}
