package publisher

import "fmt"

type Publisher struct {
	something interface{}
}

func (p *Publisher) Publish() {
	fmt.Println("vim-go")
}
