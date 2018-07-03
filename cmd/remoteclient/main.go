package main

import (
	"fmt"

	"github.com/pebbe/zmq4"
)

func main() {
	fmt.Println("vim-go")
	_ = zmq4.NewPoller()
}
