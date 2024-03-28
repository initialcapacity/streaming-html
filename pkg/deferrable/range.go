package deferrable

import "net/http"

type Range[T any] struct {
	writer  http.ResponseWriter
	channel chan T
}

func NewRange[T any](writer http.ResponseWriter, channel chan T) Range[T] {
	return Range[T]{
		writer:  writer,
		channel: channel,
	}
}

func (d Range[T]) Get() chan T {
	_ = Flush(d.writer)
	flushChannel := make(chan T)

	go func() {
		for val := range d.channel {
			_ = Flush(d.writer)
			flushChannel <- val
		}
		_ = Flush(d.writer)
		close(flushChannel)
	}()

	return flushChannel
}
