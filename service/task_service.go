package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rizkyfauziilmi/task-cli-go/model"
	"github.com/rizkyfauziilmi/task-cli-go/storage"
)

func getNextID(tasks []model.Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func findTaskIndex(tasks []model.Task, id int) int {
	for i, t := range tasks {
		if t.ID == id {
			return i
		}
	}
	return -1
}

func AddTask(tasks []model.Task, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: add <task>")
		return
	}

	desc := strings.Join(args, " ")

	task := model.Task{
		ID:          getNextID(tasks),
		Description: desc,
		Status:      model.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	fmt.Printf("Task added (ID: %d)\n", task.ID)
}

func UpdateTask(tasks []model.Task, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: update <id> <description>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	idx := findTaskIndex(tasks, id)
	if idx == -1 {
		fmt.Println("Task not found")
		return
	}

	tasks[idx].Description = strings.Join(args[1:], " ")
	tasks[idx].UpdatedAt = time.Now()

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving:", err)
	}

	fmt.Printf("Task %d updated\n", id)
}

func DeleteTask(tasks []model.Task, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: delete <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	idx := findTaskIndex(tasks, id)
	if idx == -1 {
		fmt.Println("Task not found")
		return
	}

	tasks = append(tasks[:idx], tasks[idx+1:]...)

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving:", err)
	}

	fmt.Println("Task deleted")
}

func MarkStatus(tasks []model.Task, args []string, status string) {
	if len(args) < 1 {
		fmt.Println("Usage: mark-in-progress/mark-done <id>")
		return
	}

	if status != model.StatusInProgress && status != model.StatusDone {
		fmt.Println("Invalid status")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	idx := findTaskIndex(tasks, id)
	if idx == -1 {
		fmt.Println("Task not found")
		return
	}

	tasks[idx].Status = status
	tasks[idx].UpdatedAt = time.Now()

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving:", err)
	}

	fmt.Printf("Task %d updated to %s\n", id, status)
}

func ListTasks(tasks []model.Task, args []string) {
	var filter string
	if len(args) > 0 {
		filter = strings.ToLower(args[0])
	}

	for _, t := range tasks {
		if filter != "" && t.Status != filter {
			continue
		}

		fmt.Printf("[%d] %s\nStatus: %s\nCreated: %s\nUpdated: %s\n\n",
			t.ID,
			t.Description,
			t.Status,
			t.CreatedAt.Format("2006-01-02 15:04"),
			t.UpdatedAt.Format("2006-01-02 15:04"),
		)
	}
}
