package aggregator

import "fmt"

type Aggregator struct {
	something interface{}
}

func (a *Aggregator) Connect() {
	fmt.Println("vim-go")
}
