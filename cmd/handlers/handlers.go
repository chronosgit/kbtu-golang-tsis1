package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	interfaces "github.com/chronosgit/kbtu-golang-tsis1/api/cmd/interfaces"
	json_fetch "github.com/chronosgit/kbtu-golang-tsis1/api/cmd/json_fetch"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Sophisticated Leetcode API is up and running")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var response interfaces.ITasksResponse

	tasks := json_fetch.GetTasksFromJson()

	response.Tasks = tasks

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var response interfaces.ITaskResponse
	tasks := json_fetch.GetTasksFromJson()

	var foundTask interfaces.ITask
	for _, task := range tasks {
		if task.Id == id {
			foundTask = task
			break
		}
	}

	response.Task = foundTask

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "Sophisticated Leetcode API is up and running")
	w.Write(jsonResponse)
}
