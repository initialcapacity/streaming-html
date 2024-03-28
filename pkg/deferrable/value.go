package deferrable

import "net/http"

type Value[T any] struct {
	writer  http.ResponseWriter
	channel chan T
}

func NewValue[T any](writer http.ResponseWriter, channel chan T) Value[T] {
	return Value[T]{
		writer:  writer,
		channel: channel,
	}
}

func (d Value[T]) Get() T {
	_ = Flush(d.writer)
	return <-d.channel
}
