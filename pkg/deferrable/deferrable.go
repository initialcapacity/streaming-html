package deferrable

import "net/http"

type Defer[T any] struct {
	writer http.ResponseWriter
	value  T
}

func New[T any](writer http.ResponseWriter, value T) Defer[T] {
	return Defer[T]{
		writer: writer,
		value:  value,
	}
}

func (d Defer[T]) Value() T {
	_ = http.NewResponseController(d.writer).Flush()
	return d.value
}
