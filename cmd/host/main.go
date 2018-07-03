package main

import (
	"fmt"

	"github.com/muka/go-bluetooth/bluez"
	"github.com/paypal/gatt"
	"github.com/pebbe/zmq4"
)

func main() {
	fmt.Println("vim-go")
	_ = zmq4.NewPoller()
	_ = gatt.CharNotify
	_ = bluez.FlagCharacteristicRead
}
