package messages_test

import (
	"github.com/initialcapacity/go-streaming/pkg/messages"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider_FetchAll(t *testing.T) {
	provider := messages.NewProvider(0)

	result := provider.FetchAll()
	assert.Equal(t,
		[]string{"Here's some slow content.", "It took a while to load.", "And didn't use any javascript."},
		result,
	)
}
