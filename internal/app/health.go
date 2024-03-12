package app

import "net/http"

func Health(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("{\"status\": \"up\"}"))
}
