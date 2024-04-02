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

func Index(addArtificialDelay bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(chan []string)

		go func() {
			if addArtificialDelay {
				time.Sleep(1 * time.Second)
			}
			data <- []string{
				"Here's some slower content.",
				"It took a while to load.",
				"And didn't use any javascript.",
			}
			close(data)
		}()

		_ = websupport.Render(w, Resources, "index", model{Message: deferrable.New(w, data)})
	}
}
