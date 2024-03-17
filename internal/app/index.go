package app

import (
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"net/http"
	"time"
)

func Index(addArtificialDelay bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := make(chan []string)

		go func() {
			if addArtificialDelay {
				time.Sleep(2 * time.Second)
			}
			items := []string{
				"Here's some slower content.",
				"It took a while to load.",
				"And didn't use any javascript.",
			}
			c <- items
			websupport.Flush(w)
			close(c)
		}()

		go func() {
			time.Sleep(10 * time.Millisecond)
			websupport.Flush(w)
		}()

		websupport.Render(w, Resources, "index", c)
	}
}
