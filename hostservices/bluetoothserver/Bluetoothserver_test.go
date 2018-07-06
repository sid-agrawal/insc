package bluetoothserver

import "testing"

func TestRun(t *testing.T) {
	bs := &BTSImpl{}
	bs.Run()
}

func TestPull(t *testing.T) {
	bs := &BTSImpl{}
	bs.Pull()
}
