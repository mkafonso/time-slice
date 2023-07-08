package customErrors_test

import (
	"testing"

	customErrors "github.com/mkafonso/time-slice/internal/custom-errors"
)

func TestErrorVariables(t *testing.T) {
	if customErrors.ErrorTaskNotFound.Error() != "task not found" {
		t.Errorf("ErrorTaskNotFound has an incorrect value")
	}
}
