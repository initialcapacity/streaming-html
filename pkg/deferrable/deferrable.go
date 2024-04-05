package deferrable

import "net/http"

type Deferrable[T any] struct {
	writer  http.ResponseWriter
	channel chan T
}

func New[T any](writer http.ResponseWriter, channel chan T) Deferrable[T] {
	return Deferrable[T]{
		writer:  writer,
		channel: channel,
	}
}

func (d Deferrable[T]) GetOne() T {
	d.flush()
	return <-d.channel
}

func (d Deferrable[T]) GetAll() chan T {
	d.flush()
	flushChannel := make(chan T)

	go func() {
		for val := range d.channel {
			d.flush()
			flushChannel <- val
		}
		d.flush()
		close(flushChannel)
	}()

	return flushChannel
}

func (d Deferrable[T]) flush() {
	_ = http.NewResponseController(d.writer).Flush()
}
