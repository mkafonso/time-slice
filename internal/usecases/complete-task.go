package usecases

import (
	"errors"
	"time"

	"github.com/mkafonso/time-slice/internal/models"
)

type CompleteTaskManager struct {
	Tasks models.Tasks
}

var ErrorTaskNotFound = errors.New("task not found")

func (tm *CompleteTaskManager) CompleteTask(index int) error {
	for i := range tm.Tasks {
		if i == index {
			tm.Tasks[i].Done = true
			tm.Tasks[i].CompletedAt = time.Now()
			return nil
		}
	}

	return ErrorTaskNotFound
}
