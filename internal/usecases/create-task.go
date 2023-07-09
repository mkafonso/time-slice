package usecases

import (
	"time"

	"github.com/mkafonso/time-slice/internal/models"
)

type CreateTaskManager struct {
	Tasks *models.Tasks
}

func (tm *CreateTaskManager) CreateTask(tasks interface{}, name string) {
	tasksSlice := tasks.(*models.Tasks)

	newTask := models.Task{
		Name:        name,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	*tasksSlice = append(*tasksSlice, newTask)
}
