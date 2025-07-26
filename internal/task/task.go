package task

import (
	"ToDoApp/internal/models"
	"ToDoApp/internal/storage"
	"ToDoApp/pkg/utils"
	"fmt"
)

func ReadData() {
	storage.ReadTasks()
}

func AddTask(title string) error {
	tasks, _ := storage.LoadTasks()

	newTask := models.Task{
		ID:     utils.GenerateID(),
		Title:  title,
		Status: "pending",
	}

	tasks = append(tasks, newTask)
	return storage.SaveTasks(tasks)
}

func UpdateTaskStatus(id int, status string) {
	tasks, _ := storage.LoadTasks()
	updated := false

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			updated = true
			break
		}
	}

	if updated {
		storage.SaveTasks(tasks)
		fmt.Println("Task updated.")
	} else {
		fmt.Println("Task ID not found.")
	}
}

func DeleteTask(id int) {
	tasks, _ := storage.LoadTasks()
	newTasks := []models.Task{}
	deleted := false

	for _, t := range tasks {
		if t.ID == id {
			deleted = true
			continue
		}
		newTasks = append(newTasks, t)
	}

	if deleted {
		storage.SaveTasks(newTasks)
		fmt.Println("Task deleted.")
	} else {
		fmt.Println("Task ID not found.")
	}
}
