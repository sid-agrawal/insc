package publisher

import "fmt"

// Publisher is used to ...
type Publisher struct {
	something interface{}
}

// Publish is use to PUB data.
func (p *Publisher) Publish() {
	fmt.Println("vim-go")
}
