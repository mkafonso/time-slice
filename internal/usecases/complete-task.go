package usecases

import (
	"time"

	"github.com/mkafonso/time-slice/internal/helpers"
	"github.com/mkafonso/time-slice/internal/models"
)

type CompleteTaskManager struct {
	Tasks *models.Tasks
}

func (tm *CompleteTaskManager) CompleteTask(tasks interface{}, index int) error {
	tasksSlice := tasks.(*models.Tasks)

	if index <= 0 || index > len(*tasksSlice) {
		return helpers.ErrorTaskNotFound
	}

	(*tasksSlice)[index-1].Done = true
	(*tasksSlice)[index-1].CompletedAt = time.Now()

	return nil
}
