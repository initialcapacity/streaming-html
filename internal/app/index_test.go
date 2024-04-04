package app_test

import (
	"github.com/initialcapacity/go-streaming/internal/app"
	"github.com/initialcapacity/go-streaming/pkg/testsupport"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

type controlledProvider struct {
	MessageChannel chan []string
}

func (p controlledProvider) FetchAll() []string {
	return <-p.MessageChannel
}

func TestIndex(t *testing.T) {
	messageChannel := make(chan []string)
	address, server := testsupport.StartTestServer(t, app.Handlers(controlledProvider{MessageChannel: messageChannel}))
	defer testsupport.StopTestServer(t, server)

	response, err := http.Get(address)
	assert.NoError(t, err)

	initialRead := readString(t, response.Body)
	assert.Contains(t, initialRead, "Streaming HTML")
	assert.NotContains(t, initialRead, "Success!")

	messageChannel <- []string{"some message"}

	finalRead, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(finalRead), "Success!")
	assert.Contains(t, string(finalRead), "some message")
}

func readString(t *testing.T, reader io.Reader) string {
	firstRead := make([]byte, 0, 4096)
	n, err := reader.Read(firstRead[0:4096])
	assert.NoError(t, err)
	return string(firstRead[:len(firstRead)+n])
}
