package app

import (
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"net/http"
	"time"
)

func Index(addArtificialDelay bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		websupport.Render(w, Resources, "index", nil)
		websupport.Flush(w)

		if addArtificialDelay {
			time.Sleep(2 * time.Second)
		}

		websupport.Render(w, Resources, "content", nil)
		websupport.Render(w, Resources, "end", nil)
		websupport.Flush(w)
	}
}
