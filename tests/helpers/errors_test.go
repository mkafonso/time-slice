package helpers_test

import (
	"testing"

	"github.com/mkafonso/time-slice/internal/helpers"
)

func TestErrorVariables(t *testing.T) {
	if helpers.ErrorTaskNotFound.Error() != "task not found" {
		t.Errorf("ErrorTaskNotFound has an incorrect value")
	}
}
