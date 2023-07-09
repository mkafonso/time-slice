package usecases

import (
	"github.com/mkafonso/time-slice/internal/helpers"
	"github.com/mkafonso/time-slice/internal/models"
)

type DeleteTaskManager struct {
	Tasks *models.Tasks
}

func (tm *DeleteTaskManager) DeleteTask(tasks interface{}, index int) error {
	tasksSlice := tasks.(*models.Tasks)

	if index <= 0 || index > len(*tasksSlice) {
		return helpers.ErrorTaskNotFound
	}

	*tasksSlice = append((*tasksSlice)[:index-1], (*tasksSlice)[index:]...)

	return nil
}
