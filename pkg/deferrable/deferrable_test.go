package deferrable_test

import (
	"github.com/initialcapacity/go-streaming/pkg/deferrable"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDefer(t *testing.T) {
	writer := ResponseWriterSpy{}
	d := deferrable.New(&writer, "pickles")

	value := d.Value()

	assert.Equal(t, "pickles", value)
	assert.True(t, writer.FlushCalled)
}

type ResponseWriterSpy struct {
	FlushCalled bool
}

func (w *ResponseWriterSpy) Header() http.Header {
	return map[string][]string{}
}

func (w *ResponseWriterSpy) Write([]byte) (int, error) {
	return 0, nil
}

func (w *ResponseWriterSpy) WriteHeader(statusCode int) {
}

func (w *ResponseWriterSpy) Flush() {
	w.FlushCalled = true
}
