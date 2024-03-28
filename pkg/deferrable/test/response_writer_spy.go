package deferrable_test

import (
	"net/http"
	"sync/atomic"
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
