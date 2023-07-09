package usecases

import (
	"time"

	"github.com/mkafonso/time-slice/internal/models"
)

type CreateTaskManager struct {
	Tasks *models.Tasks
}

func (tm *CreateTaskManager) CreateTask(name string) {
	newTask := models.Task{
		Name:        name,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	*tm.Tasks = append(*tm.Tasks, newTask)
}
