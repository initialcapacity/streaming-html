package websupport

import (
	"os"
	"strconv"
)

func IntegerEnvironmentVariable(variableName string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(variableName))
	if err != nil {
		return defaultValue
	}

	return value
}
