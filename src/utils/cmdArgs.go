package utils

import (
	"errors"
	"os"
)

func GetCmdArgs(field string) (string, error) {
	ind := -1

	for i, val := range os.Args {
		if val == field {
			ind = i + 1
			break
		}
	}

	if ind >= len(os.Args) || ind == -1 {
		return "", errors.New("argument not found for " + field)
	}

	return os.Args[ind], nil
}
