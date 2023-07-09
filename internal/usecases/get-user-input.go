package usecases

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func GetUserInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(scanner.Text()) == 0 {
		return "", errors.New("Empty task")
	}

	return scanner.Text(), nil
}
