package main

import (
	"github.com/initialcapacity/go-streaming/internal/app"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"log"
)

func main() {
	host := websupport.EnvironmentVariable("HOST", "")
	port := websupport.EnvironmentVariable("PORT", 8777)

	server := websupport.NewServer(app.Handlers(true))

	_, done := server.Start(host, port)
	log.Fatal(<-done)
}
