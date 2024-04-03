package app

import (
	"github.com/initialcapacity/go-streaming/pkg/deferrable"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"net/http"
	"time"
)

type model struct {
	Message deferrable.Deferrable[[]string]
}

type waiter interface {
	Wait() <-chan time.Time
}

func Index(delay waiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(chan []string)

		go func() {
			<-delay.Wait()
			data <- []string{"Here's some slow content.", "It took a while to load.", "And didn't use any javascript."}
			close(data)
		}()

		_ = websupport.Render(w, Resources, "index", model{Message: deferrable.New(w, data)})
	}
}
