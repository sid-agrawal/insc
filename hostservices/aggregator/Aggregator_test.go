package aggregator

import "testing"

func TestConnectOk(t *testing.T) {
	a := &Aggregator{}
	a.Connect()
}
