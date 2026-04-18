package storage

import (
	"encoding/json"
	"os"

	"github.com/rizkyfauziilmi/task-cli-go/model"
)

const filePath = "tasks.json"

func EnsureFile() error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return os.WriteFile(filePath, []byte("[]"), 0644)
	}
	return nil
}

func LoadTasks() ([]model.Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	if len(data) > 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

func SaveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
