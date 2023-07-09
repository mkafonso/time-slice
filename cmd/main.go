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
	deleteTaskFlag := flag.Int("delete", 0, "Delete a task")
	listAllTasksFlag := flag.Bool("list", false, "List all the tasks")
	flag.Parse()

	// create a pointer to the tasks slice
	tasks := &models.Tasks{}

	// create the task managers
	taskManagerStore := &usecases.StoreTasksInFileManager{Tasks: tasks}
	taskManagerCreate := &usecases.CreateTaskManager{Tasks: tasks}
	taskManagerComplete := &usecases.CompleteTaskManager{Tasks: tasks}
	taskManagerDelete := &usecases.DeleteTaskManager{Tasks: tasks}
	taskManagerList := &usecases.ListAllTasksManager{Tasks: tasks}

	// load tasks from file
	tasksFromFile, _ := usecases.LoadTasksFromFile("tasks.json")

	switch {
	case *listAllTasksFlag:
		taskManagerList.ListAllTasks(tasksFromFile)

	case *addTaskFlag:
		task, err := usecases.GetUserInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		taskManagerCreate.CreateTask(tasksFromFile, task)

		error := taskManagerStore.StoreTasksInFile(tasksFromFile, "tasks.json")
		if error != nil {
			fmt.Fprintln(os.Stderr, error.Error())
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

	case *deleteTaskFlag >= 0:
		err := taskManagerDelete.DeleteTask(tasksFromFile, *deleteTaskFlag)
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
