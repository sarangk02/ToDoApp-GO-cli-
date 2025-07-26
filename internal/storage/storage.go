package storage

import (
	"ToDoApp/internal/models"
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "tasks.json"

func ReadTasks() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("\nYour Tasks:")
	for _, t := range tasks {
		fmt.Printf("ID: %d | %s | Status: %s\n", t.ID, t.Title, t.Status)
	}
}

func SaveTasks(tasks []models.Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // pretty print
	return encoder.Encode(tasks)
}

func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task

	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil // no file = empty list
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}
	return tasks, nil
}
