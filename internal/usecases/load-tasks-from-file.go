package usecases

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/mkafonso/time-slice/internal/models"
)

func LoadTasksFromFile(filename string) (*models.Tasks, error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &models.Tasks{}, nil
		}
		return &models.Tasks{}, err
	}

	if len(file) == 0 {
		return &models.Tasks{}, nil
	}

	var tasks models.Tasks
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return &models.Tasks{}, nil
	}

	return &tasks, nil
}
