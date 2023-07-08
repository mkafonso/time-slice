package usecases_test

import (
	"testing"

	"github.com/mkafonso/time-slice/internal/usecases"
)

func TestUseCase_CreateTask(t *testing.T) {
	taskManager := usecases.CreateTaskManager{}

	taskManager.CreateTask("New Task")

	// check length of taskManager.Tasks
	if len(taskManager.Tasks) != 1 {
		t.Errorf("Expected list to have length 1. Got %d", len(taskManager.Tasks))
	}

	// check if the item was added correctly
	task := taskManager.Tasks[0]
	if task.Name != "New Task" {
		t.Errorf("Incorrect task name. Expected 'New Task' but got '%s'", task.Name)
	}

	// check task status
	if task.Done != false {
		t.Errorf("Incorrect status. Expected 'False' but go 'true'")
	}

	// created_at must be valid
	if task.CreatedAt.IsZero() {
		t.Error("The field 'CreatedAt' is incorrect")
	}

	// completed_at must be valid
	if task.CompletedAt.IsZero() {
		t.Error("The field 'CompletedAt' is incorrect")
	}
}

func TestUseCase_CreateTaskWithMultipleTasks(t *testing.T) {
	taskManager := usecases.CreateTaskManager{}

	// add three tasks
	taskManager.CreateTask("Task 1")
	taskManager.CreateTask("Task 2")
	taskManager.CreateTask("Task 3")

	// check taskManager.Tasks length
	if len(taskManager.Tasks) != 3 {
		t.Errorf("Expected list to have length 3. Got %d", len(taskManager.Tasks))
	}

	// check if the tasks were added correctly
	expectedNames := []string{"Task 1", "Task 2", "Task 3"}
	for i, task := range taskManager.Tasks {
		if task.Name != expectedNames[i] {
			t.Errorf("Incorrect task name. Expected '%s', but got '%s'", expectedNames[i], task.Name)
		}
	}
}
