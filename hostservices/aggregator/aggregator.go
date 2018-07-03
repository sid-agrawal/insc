package aggregator

import "fmt"

// Aggregator implements ...
type Aggregator struct {
	something interface{}
}

// Connect to something
func (a *Aggregator) Connect() {
	fmt.Println("vim-go")
}
