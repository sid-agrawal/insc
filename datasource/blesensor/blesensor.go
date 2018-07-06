package blesensor

import "fmt"

// BLESensor implements the datasource interface.
type BLESensor struct {
	something interface{}
}

// Connect is used to connect to the Aggregator.
func (b *BLESensor) Connect() {
	fmt.Println("vim-go")
}
