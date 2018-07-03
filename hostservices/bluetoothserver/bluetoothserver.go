package bluetoothserver

import "fmt"

type BluetoothServer struct {
	something interface{}
}

func (bs *BluetoothServer) Connect() {
	fmt.Println("vim-go")
}
