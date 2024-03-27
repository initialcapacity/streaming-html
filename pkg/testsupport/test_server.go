package testsupport

import (
	"fmt"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"net/http"
	"testing"
	"time"
)

func StartTestServer(t *testing.T, handlers websupport.Handlers) (string, *websupport.Server) {
	server := websupport.NewServer(handlers)

	port, _ := server.Start("localhost", 0)
	AssertHealthy(t, port, "/health")

	return fmt.Sprintf("http://localhost:%d", port), server
}

func StopTestServer(t *testing.T, server *websupport.Server) {
	err := server.Stop()
	if err != nil {
		t.Errorf("unable to stop server: %s", err)
	}
}

func AssertHealthy(t *testing.T, port int, path string) {
	statusCode := make(chan int)

	go func() {
		for {
			resp, err := http.Get(fmt.Sprintf("http://localhost:%d%s", port, path))
			if err == nil {
				statusCode <- resp.StatusCode
				return
			}
		}
	}()

	select {
	case code := <-statusCode:
		if code != http.StatusOK {
			t.Errorf("server responded with a non 200 code: %d", code)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("server did not respond in 100 milliseconds")
	}
}
