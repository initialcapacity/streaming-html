package deferrable_test

import (
	"github.com/initialcapacity/go-streaming/pkg/deferrable"
	"github.com/initialcapacity/go-streaming/pkg/deferrable/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannel(t *testing.T) {
	writer := deferrable_test.ResponseWriterSpy{}
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "pickles"
	}()

	value := deferrable.NewValue(&writer, channel).Get()

	assert.Equal(t, "pickles", value)
	assert.Equal(t, uint64(1), writer.FlushCalls.Load())
}
