package remoteclient

import "testing"

func TestConnectOk(t *testing.T) {
	r := &RemoteClient{}
	r.Connect()
}
