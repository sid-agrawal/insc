package bluetoothserver

import "testing"

func TestConnectOk(t *testing.T) {
	bs := &BluetoothServer{}
	bs.Connect()
}
