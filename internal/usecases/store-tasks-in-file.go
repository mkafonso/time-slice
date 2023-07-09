package usecases

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mkafonso/time-slice/internal/models"
)

type StoreTasksInFileManager struct {
	Tasks *models.Tasks
}

func (tm *StoreTasksInFileManager) StoreTasksInFile(tasks interface{}, filename string) error {
	tasksSlice := tasks.(*models.Tasks)

	data, err := json.MarshalIndent(*tasksSlice, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
