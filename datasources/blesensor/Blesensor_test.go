package blesensor

import "testing"

func TestConnectOk(t *testing.T) {
	b := &BleSensor{}
	b.Connect()
}
