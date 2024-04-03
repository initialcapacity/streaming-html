package main

import (
	"github.com/initialcapacity/go-streaming/internal/app"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"log"
	"time"
)

type oneSecondWaiter struct {
}

func (w oneSecondWaiter) Wait() <-chan time.Time {
	return time.After(1 * time.Second)
}

func main() {
	host := websupport.EnvironmentVariable("HOST", "")
	port := websupport.EnvironmentVariable("PORT", 8777)

	server := websupport.NewServer(app.Handlers(oneSecondWaiter{}))

	_, done := server.Start(host, port)
	log.Fatal(<-done)
}
