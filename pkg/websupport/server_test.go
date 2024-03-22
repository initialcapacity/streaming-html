package websupport_test

import (
	"fmt"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	server := websupport.NewServer(func(mux *http.ServeMux) {
		mux.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)
			_, err := writer.Write([]byte("You passed the test"))
			assert.NoError(t, err)
		})
	})

	port, _ := server.Start("localhost", 0)
	err := server.WaitUntilHealthy("/")
	assert.NoError(t, err)
	defer func(server *websupport.Server) {
		_ = server.Stop()
	}(server)

	response, err := http.Get(fmt.Sprintf("http://localhost:%d/", port))
	assert.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), "You passed the test")
}
