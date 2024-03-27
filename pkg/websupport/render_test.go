package websupport_test

import (
	"bytes"
	"github.com/initialcapacity/go-streaming/pkg/websupport"
	"github.com/initialcapacity/go-streaming/pkg/websupport/test"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

type TestData struct {
	Message string
}

func TestRender(t *testing.T) {
	writer := httptest.ResponseRecorder{Body: new(bytes.Buffer)}
	data := TestData{Message: "Hello"}

	err := websupport.Render(&writer, websupport_test.Resources, "test", data)
	assert.NoError(t, err)

	body, _ := io.ReadAll(writer.Body)
	assert.Equal(t, `
    <html lang="en">
    <body>
    <h1>Hello</h1>
    </body>
    </html>
`, string(body))
}
