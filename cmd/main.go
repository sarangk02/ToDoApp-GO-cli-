package main

import (
	"ToDoApp/internal/task"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"` // "pending" or "done"
}

func takeAction(choice int) {
	reader := bufio.NewReader(os.Stdin)

	switch choice {
	case 1:
		task.ReadData()
	case 2:
		fmt.Print("Enter task title: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)
		if err := task.AddTask(title); err == nil {
			fmt.Println("Task added.")
		}
	case 3:
		fmt.Print("Enter task ID to update: ")
		var id int
		fmt.Scanln(&id)
		fmt.Print("Enter new status (pending/done): ")
		status, _ := reader.ReadString('\n')
		status = strings.TrimSpace(status)
		task.UpdateTaskStatus(id, status)
	case 4:
		fmt.Print("Enter task ID to delete: ")
		var id int
		fmt.Scanln(&id)
		task.DeleteTask(id)
	case 5:
		fmt.Println("Exiting.")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice.")
	}
}

func main() {
	for {
		fmt.Println("\n--- To-Do App ---")
		fmt.Println("1. Read Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Update Task Status")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		var inp int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&inp)
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		takeAction(inp)
	}
}
