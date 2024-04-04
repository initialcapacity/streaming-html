package app_test

import (
	"github.com/initialcapacity/go-streaming/internal/app"
	"github.com/initialcapacity/go-streaming/pkg/messages"
	"github.com/initialcapacity/go-streaming/pkg/testsupport"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"testing"
)

func TestHealth(t *testing.T) {
	server := websupport.NewServer(app.Handlers(messages.NewProvider(0)))
	port, _ := server.Start("localhost", 0)
	testsupport.AssertHealthy(t, port, "/health")
}
