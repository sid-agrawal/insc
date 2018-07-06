package bluetoothserver

import "fmt"

// BluetoothServer is the interface for a BT Server
type BluetoothServer interface {
	Run()
	Pull(string) (string, error)
}

// BTSImpl implements ...
type BTSImpl struct {
	something interface{}
}

// Run is used to start the BT Server
func (bs *BTSImpl) Run() {
	fmt.Println("vim-go")
}

// Pull is used to Pull source from a BT device
func (bs *BTSImpl) Pull() {
	fmt.Println("vim-go")
}
