package blesensor

import "testing"

func TestConnectOk(t *testing.T) {
	b := &BLESensor{}
	b.Connect()
}
