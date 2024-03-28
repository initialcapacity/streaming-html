package deferrable

import "net/http"

func Flush(w http.ResponseWriter) error {
	return http.NewResponseController(w).Flush()
}
