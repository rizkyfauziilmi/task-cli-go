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
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []model.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		// return empty slice if file empty
		return []model.Task{}, nil
	}

	return tasks, nil
}

func SaveTasks(tasks []model.Task) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(tasks)
}
