package deferrable_test

import (
	"github.com/initialcapacity/go-streaming/pkg/deferrable"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync/atomic"
	"testing"
)

type ResponseWriterSpy struct {
	FlushCalls atomic.Uint64
}

func (w *ResponseWriterSpy) Header() http.Header {
	return map[string][]string{}
}

func (w *ResponseWriterSpy) Write([]byte) (int, error) {
	return 0, nil
}

func (w *ResponseWriterSpy) WriteHeader(_ int) {
}

func (w *ResponseWriterSpy) Flush() {
	w.FlushCalls.Add(1)
}

func TestGetOne(t *testing.T) {
	writer := ResponseWriterSpy{}
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "pickles"
	}()

	value := deferrable.New(&writer, channel).GetOne()

	assert.Equal(t, "pickles", value)
	assert.Equal(t, uint64(1), writer.FlushCalls.Load())
}

func TestGetAll(t *testing.T) {
	writer := ResponseWriterSpy{}
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "pickles"
		channel <- "chicken"
	}()

	values := deferrable.New(&writer, channel).GetAll()

	assert.Equal(t, uint64(1), writer.FlushCalls.Load())
	assert.Equal(t, "pickles", <-values)
	assert.Equal(t, "chicken", <-values)
	assert.Equal(t, uint64(3), writer.FlushCalls.Load())
}
