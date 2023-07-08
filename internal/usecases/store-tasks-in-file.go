package usecases

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mkafonso/time-slice/internal/models"
)

type StoreTasksInFileManager struct {
	Tasks models.Tasks
}

func (tm *StoreTasksInFileManager) StoreTasksInFile(filename string) error {
	data, err := json.Marshal(tm.Tasks)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0666)
}
