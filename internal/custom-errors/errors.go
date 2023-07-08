package customErrors

import "errors"

var (
	ErrorTaskNotFound = errors.New("task not found")
)
