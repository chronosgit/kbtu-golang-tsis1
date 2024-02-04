package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	leetcode "github.com/chronosgit/kbtu-golang-tsis1/api/pkg/leetcode"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Sophisticated Leetcode API is up and running")
}

func Tasks(w http.ResponseWriter, r *http.Request) {
	var response ITasksResponse

	tasks := leetcode.GetTasksFromJson()

	response.Tasks = tasks

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func TaskById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var response ITaskResponse
	tasks := leetcode.GetTasksFromJson()

	var foundTask leetcode.ITask
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
