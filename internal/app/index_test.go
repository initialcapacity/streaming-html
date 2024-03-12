package app_test

import (
	"github.com/initialcapacity/go-streaming/internal/app"
	"github.com/initialcapacity/go-streaming/pkg/testsupport"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestIndex(t *testing.T) {
	address, server := testsupport.StartTestServer(t, app.Handlers(false))
	defer testsupport.StopTestServer(t, server)

	response, err := http.Get(address)
	assert.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), "Go streaming")
	assert.Contains(t, string(body), "Success!")
}
