package websupport

import (
	"errors"
	"net/http"
)

type FlushWriter struct {
	writer     http.ResponseWriter
	controller *http.ResponseController
}

func NewFlushWriter(writer http.ResponseWriter) *FlushWriter {
	return &FlushWriter{
		writer:     writer,
		controller: http.NewResponseController(writer),
	}
}

func (fw FlushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.writer.Write(p)
	flushErr := fw.controller.Flush()
	return n, errors.Join(err, flushErr)
}
