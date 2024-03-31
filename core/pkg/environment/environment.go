package environment

import (
	"fmt"
	"os"
)

func GetWorkingDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Failed to get current working directory: %w", err)
	}
	return cwd, nil
}
