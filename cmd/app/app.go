package main

import (
	"github.com/initialcapacity/go-streaming/internal/app"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"log"
)

func main() {
	port := websupport.IntegerEnvironmentVariable("PORT", 8777)

	server := websupport.NewServer(app.Handlers(true))

	_, done := server.Start(port)
	log.Fatal(<-done)
}
