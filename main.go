package main

import (
	"net/http"

	handlers "github.com/chronosgit/kbtu-golang-tsis1/api/cmd/handlers"

	"github.com/gorilla/mux"
)

func initRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", handlers.GetTaskById).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}

func main() {
	initRouter()
}
