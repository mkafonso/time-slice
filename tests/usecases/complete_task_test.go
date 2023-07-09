package usecases_test

import (
	"testing"
	"time"

	customErrors "github.com/mkafonso/time-slice/internal/custom-errors"
	"github.com/mkafonso/time-slice/internal/models"
	"github.com/mkafonso/time-slice/internal/usecases"
)

func TestUseCase_CompleteTask(t *testing.T) {
	// create a task manager with some tasks
	tasks := &models.Tasks{
		{Name: "Task 1", Done: false, CreatedAt: time.Now()},
		{Name: "Task 2", Done: false, CreatedAt: time.Now()},
		{Name: "Task 3", Done: false, CreatedAt: time.Now()},
	}

	taskManager := &usecases.CompleteTaskManager{Tasks: tasks}

	// complete the task with index 1
	err := taskManager.CompleteTask(tasks, 2)
	if err != nil {
		t.Errorf("Error completing task: %v", err)
	}

	// verify that the task's Done status is now true and CompletedAt is set
	task := (*tasks)[1]
	if !task.Done {
		t.Errorf("Task 1 should be marked as done")
	}
	if task.CompletedAt.IsZero() {
		t.Errorf("Task 1's CompletedAt time should not be zero")
	}
}

func TestUseCase_CompleteTask_NotFound(t *testing.T) {
	// create a task manager with some tasks
	tasks := &models.Tasks{
		{Name: "Task 1", Done: false, CreatedAt: time.Now()},
		{Name: "Task 2", Done: false, CreatedAt: time.Now()},
		{Name: "Task 3", Done: false, CreatedAt: time.Now()},
	}

	taskManager := &usecases.CompleteTaskManager{Tasks: tasks}

	// complete a non-existent task
	err := taskManager.CompleteTask(tasks, 10)
	if err != customErrors.ErrorTaskNotFound {
		t.Errorf("Expected ErrorTaskNotFound, but got %v", err)
	}
}
