package leetcode

import (
	"encoding/json"
	"log"
	"os"
)

type ITask struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	Link       string `json:"link"`
}

func GetTasksFromJson() []ITask {
	jsonFilePath := "./tasks.json"

	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []ITask

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	return tasks
}
