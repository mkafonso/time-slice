package helpers

import (
	"fmt"
)

const (
	ColorDefault = "\x1b[39m"
	ColorPending = "\x1b[90m"
	ColorSuccess = "\x1b[93m"
)

func PendingColor(s string) string {
	return fmt.Sprintf("%s%s%s", ColorPending, s, ColorDefault)
}

func DoneColor(s string) string {
	return fmt.Sprintf("%s%s%s", ColorSuccess, s, ColorDefault)
}
