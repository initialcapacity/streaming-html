package testsupport

import (
	"fmt"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"testing"
)

func StartTestServer(t *testing.T, handlers websupport.Handlers) (string, *websupport.Server) {
	server := websupport.NewServer(handlers)

	port, _ := server.Start("localhost", 0)
	err := server.WaitUntilHealthy("/health")
	if err != nil {
		t.Errorf("unable to start server: %s", err)
	}

	return fmt.Sprintf("http://localhost:%d", port), server
}

func StopTestServer(t *testing.T, server *websupport.Server) {
	err := server.Stop()
	if err != nil {
		t.Errorf("unable to stop server: %s", err)
	}
}
