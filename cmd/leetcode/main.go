package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/tasks", Tasks).Methods("GET")
	router.HandleFunc("/task/{id}", TaskById).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}

func main() {
	initRouter()
}
