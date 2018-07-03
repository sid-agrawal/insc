package blesensor

import "fmt"

type BleSensor struct {
	something interface{}
}

func (b *BleSensor) Connect() {
	fmt.Println("vim-go")
}
