package publisher

import "testing"

func TestPublishOk(t *testing.T) {
	p := &Publisher{}
	p.Publish()
}
