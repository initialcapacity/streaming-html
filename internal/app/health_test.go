package app_test

import (
	"github.com/initialcapacity/go-streaming/internal/app"
	"github.com/initialcapacity/go-streaming/pkg/testsupport"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"testing"
	"time"
)

type immediateWaiter struct {
}

func (w immediateWaiter) Wait() <-chan time.Time {
	c := make(chan time.Time)
	go func() {
		c <- time.Now()
	}()
	return c
}

func TestHealth(t *testing.T) {
	server := websupport.NewServer(app.Handlers(immediateWaiter{}))
	port, _ := server.Start("localhost", 0)
	testsupport.AssertHealthy(t, port, "/health")
}
