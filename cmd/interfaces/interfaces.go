package interfaces

type ITask struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	Link       string `json:"link"`
}

type ITaskResponse struct {
	Task ITask `json:"task"`
}

type ITasksResponse struct {
	Tasks []ITask `json:"tasks"`
}
