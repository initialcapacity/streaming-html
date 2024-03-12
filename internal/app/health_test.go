package app_test

import (
	"github.com/initialcapacity/go-streaming/internal/app"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealth(t *testing.T) {
	server := websupport.NewServer(app.Handlers(false))
	_, _ = server.Start(0)
	err := server.WaitUntilHealthy("/health")
	assert.NoError(t, err)
}
