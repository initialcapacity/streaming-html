package app

import (
	"io/fs"
	"net/http"
)

func Handlers(addArtificialDelay bool) func(mux *http.ServeMux) {
	return func(mux *http.ServeMux) {
		mux.HandleFunc("GET /", Index(addArtificialDelay))
		mux.HandleFunc("GET /health", Health)

		static, _ := fs.Sub(Resources, "resources/static")
		fileServer := http.FileServer(http.FS(static))
		mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	}
}
