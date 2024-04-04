package app

import (
	"github.com/initialcapacity/go-streaming/pkg/deferrable"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"net/http"
)

type model struct {
	Message deferrable.Deferrable[[]string]
}

type messageProvider interface {
	FetchAll() []string
}

func Index(provider messageProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(chan []string)
		go func() {
			data <- provider.FetchAll()
		}()

		_ = websupport.Render(w, Resources, "index", model{Message: deferrable.New(w, data)})
	}
}
