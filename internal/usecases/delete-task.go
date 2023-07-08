package usecases

import (
	customErrors "github.com/mkafonso/time-slice/internal/custom-errors"
	"github.com/mkafonso/time-slice/internal/models"
)

type DeleteTaskManager struct {
	Tasks models.Tasks
}

func (tm *DeleteTaskManager) DeleteTask(index int) error {
	for i := range tm.Tasks {
		if i == index {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			return nil
		}
	}

	return customErrors.ErrorTaskNotFound
}
