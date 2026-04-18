package main

import (
	"fmt"
	"os"

	"github.com/rizkyfauziilmi/task-cli-go/model"
	"github.com/rizkyfauziilmi/task-cli-go/service"
	"github.com/rizkyfauziilmi/task-cli-go/storage"
)

func main() {
	if err := storage.EnsureFile(); err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		panic(err)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		service.AddTask(tasks, args)
	case "update":
		service.UpdateTask(tasks, args)
	case "delete":
		service.DeleteTask(tasks, args)
	case "mark-in-progress":
		service.MarkStatus(tasks, args, model.StatusInProgress)
	case "mark-done":
		service.MarkStatus(tasks, args, model.StatusDone)
	case "list":
		service.ListTasks(tasks, args)
	default:
		fmt.Println("Unknown command")
	}
}
