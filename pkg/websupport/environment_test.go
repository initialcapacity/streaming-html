package websupport_test

import (
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func set(t *testing.T, name, value string) {
	err := os.Setenv(name, value)
	assert.NoError(t, err)
}

func unset(t *testing.T, name string) {
	err := os.Unsetenv(name)
	assert.NoError(t, err)
}

func TestEnvironmentVariable(t *testing.T) {
	set(t, "NOT_NUMBER", "hello")
	set(t, "NUMBER", "39")
	unset(t, "NOT_PRESENT")

	assert.Equal(t, "hello", websupport.EnvironmentVariable("NOT_NUMBER", "pickles"))
	assert.Equal(t, "39", websupport.EnvironmentVariable("NUMBER", "pickles"))
	assert.Equal(t, "pickles", websupport.EnvironmentVariable("NOT_PRESENT", "pickles"))

	assert.Equal(t, 867, websupport.EnvironmentVariable("NOT_NUMBER", 867))
	assert.Equal(t, 39, websupport.EnvironmentVariable("NUMBER", 867))
	assert.Equal(t, 867, websupport.EnvironmentVariable("NOT_PRESENT", 867))
}
