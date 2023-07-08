package usecases_test

import (
	"testing"
	"time"

	customErrors "github.com/mkafonso/time-slice/internal/custom-errors"
	"github.com/mkafonso/time-slice/internal/models"
	"github.com/mkafonso/time-slice/internal/usecases"
)

func TestUseCase_DeleteTask(t *testing.T) {
	// create a task manager with some tasks
	taskManager := usecases.DeleteTaskManager{
		Tasks: models.Tasks{
			{Name: "Task 1", Done: false, CreatedAt: time.Now()},
			{Name: "Task 2", Done: false, CreatedAt: time.Now()},
			{Name: "Task 3", Done: false, CreatedAt: time.Now()},
		},
	}

	// delete the task with index 1
	err := taskManager.DeleteTask(1)
	if err != nil {
		t.Errorf("Error deleting task: %v", err)
	}

	// verify that the task with index 1 is no longer in the list of tasks
	for i := range taskManager.Tasks {
		if taskManager.Tasks[i].Name == "Task 2" {
			t.Errorf("Task 2 should have been removed, but it is still present in the list")
		}
	}
}

func TestUseCase_DeleteTask_NotFound(t *testing.T) {
	// create a task manager with some tasks
	taskManager := usecases.DeleteTaskManager{
		Tasks: models.Tasks{
			{Name: "Task 1", Done: false, CreatedAt: time.Now()},
			{Name: "Task 2", Done: false, CreatedAt: time.Now()},
			{Name: "Task 3", Done: false, CreatedAt: time.Now()},
		},
	}

	// delete a non-existent task with index 4
	err := taskManager.DeleteTask(3)
	if err != customErrors.ErrorTaskNotFound {
		t.Errorf("Expected ErrorTaskNotFound, but got %v", err)
	}
}
