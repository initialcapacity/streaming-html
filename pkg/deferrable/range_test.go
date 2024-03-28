package deferrable_test

import (
	"github.com/initialcapacity/go-streaming/pkg/deferrable"
	"github.com/initialcapacity/go-streaming/pkg/deferrable/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	writer := deferrable_test.ResponseWriterSpy{}
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "pickles"
		channel <- "chicken"
	}()

	values := deferrable.NewRange(&writer, channel).Get()

	assert.Equal(t, uint64(1), writer.FlushCalls.Load())
	assert.Equal(t, "pickles", <-values)
	assert.Equal(t, "chicken", <-values)
	assert.Equal(t, uint64(3), writer.FlushCalls.Load())
}
