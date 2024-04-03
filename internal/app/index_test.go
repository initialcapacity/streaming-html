package app_test

import (
	"github.com/initialcapacity/go-streaming/internal/app"
	"github.com/initialcapacity/go-streaming/pkg/testsupport"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

type controlledWaiter struct {
	C chan time.Time
}

func (w controlledWaiter) Wait() <-chan time.Time {
	return w.C
}

func TestIndex(t *testing.T) {
	delayChannel := make(chan time.Time)
	waiter := controlledWaiter{C: delayChannel}

	address, server := testsupport.StartTestServer(t, app.Handlers(waiter))
	defer testsupport.StopTestServer(t, server)

	response, err := http.Get(address)
	assert.NoError(t, err)

	initialRead := readString(t, response.Body)
	assert.Contains(t, initialRead, "Streaming HTML")
	assert.NotContains(t, initialRead, "Success!")

	delayChannel <- time.Now()

	finalRead, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(finalRead), "Success!")
}

func readString(t *testing.T, reader io.Reader) string {
	firstRead := make([]byte, 0, 4096)
	n, err := reader.Read(firstRead[0:4096])
	assert.NoError(t, err)
	return string(firstRead[:len(firstRead)+n])
}
