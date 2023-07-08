package usecases

import (
	"time"

	customErrors "github.com/mkafonso/time-slice/internal/custom-errors"
	"github.com/mkafonso/time-slice/internal/models"
)

type CompleteTaskManager struct {
	Tasks models.Tasks
}

func (tm *CompleteTaskManager) CompleteTask(index int) error {
	for i := range tm.Tasks {
		if i == index {
			tm.Tasks[i].Done = true
			tm.Tasks[i].CompletedAt = time.Now()
			return nil
		}
	}

	return customErrors.ErrorTaskNotFound
}
