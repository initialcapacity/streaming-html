package websupport

import (
	"os"
	"strconv"
)

func EnvironmentVariable(variableName string, defaultValue string) string {
	value, found := os.LookupEnv(variableName)
	if found {
		return value
	} else {
		return defaultValue
	}
}

func IntegerEnvironmentVariable(variableName string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(variableName))
	if err != nil {
		return defaultValue
	}

	return value
}
