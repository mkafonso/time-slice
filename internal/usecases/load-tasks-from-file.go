package usecases

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/mkafonso/time-slice/internal/models"
)

type LoadFromFileManager struct {
	Tasks models.Tasks
}

func (tm *LoadFromFileManager) LoadTasksFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	err = json.Unmarshal(file, tm)
	if err != nil {
		return nil
	}

	return nil
}
