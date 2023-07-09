package usecases

import (
	"time"

	customErrors "github.com/mkafonso/time-slice/internal/custom-errors"
	"github.com/mkafonso/time-slice/internal/models"
)

type CompleteTaskManager struct {
	Tasks *models.Tasks
}

func (tm *CompleteTaskManager) CompleteTask(tasks interface{}, index int) error {
	tasksSlice := tasks.(*models.Tasks)

	if index <= 0 || index > len(*tasksSlice) {
		return customErrors.ErrorTaskNotFound
	}

	(*tasksSlice)[index-1].Done = true
	(*tasksSlice)[index-1].CompletedAt = time.Now()

	return nil
}
