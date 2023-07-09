package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mkafonso/time-slice/internal/models"
	"github.com/mkafonso/time-slice/internal/usecases"
)

func main() {
	// Define flags
	addTaskFlag := flag.Bool("add", false, "Create a new task")
	completeTaskFlag := flag.Int("complete", 0, "Complete a task")
	flag.Parse()

	// create a pointer to the tasks slice
	tasks := &models.Tasks{}

	// create the task managers
	taskManagerStore := &usecases.StoreTasksInFileManager{Tasks: tasks}
	taskManagerCreate := &usecases.CreateTaskManager{Tasks: tasks}
	taskManagerComplete := &usecases.CompleteTaskManager{Tasks: tasks}

	// load tasks from file
	tasksFromFile, _ := usecases.LoadTasksFromFile("tasks.json")

	switch {
	case *addTaskFlag:
		taskManagerCreate.CreateTask("First task")
		taskManagerCreate.CreateTask("Second task")
		err := taskManagerStore.StoreTasksInFile(tasksFromFile, "tasks.json")
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *completeTaskFlag >= 0:
		err := taskManagerComplete.CompleteTask(tasksFromFile, *completeTaskFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = taskManagerStore.StoreTasksInFile(tasksFromFile, "tasks.json")
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stdout, "Invalid command")
		os.Exit(1)
	}
}
