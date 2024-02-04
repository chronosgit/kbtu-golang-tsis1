package main

import (
	leetcode "github.com/chronosgit/kbtu-golang-tsis1/api/pkg/leetcode"
)

type ITaskResponse struct {
	Task leetcode.ITask `json:"task"`
}

type ITasksResponse struct {
	Tasks []leetcode.ITask `json:"tasks"`
}
