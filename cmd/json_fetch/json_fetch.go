package json_fetch

import (
	"encoding/json"
	"log"
	"os"

	interfaces "github.com/chronosgit/kbtu-golang-tsis1/api/cmd/interfaces"
)

func GetTasksFromJson() []interfaces.ITask {
	jsonFilePath := "./tasks.json"

	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []interfaces.ITask

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	return tasks
}
