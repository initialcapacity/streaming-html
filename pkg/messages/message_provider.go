package messages

import "time"

type Provider struct {
	delay time.Duration
}

func NewProvider(delay time.Duration) Provider {
	return Provider{delay: delay}
}

func (provider Provider) FetchAll() []string {
	time.Sleep(provider.delay)

	return []string{"Here's some slow content.", "It took a while to load.", "And didn't use any javascript."}
}
