package publisher

import "fmt"

// Publisher is the API for the package
type Publisher interface {
	ReadFromDB()
	Publish()
}

// PImpl is used to ...
type PImpl struct {
	something interface{}
}

// Publish is use to PUB data.
func (p *PImpl) Publish() {
	fmt.Println("vim-go")
}

// ReadFromDB is use to PUB data.
func (p *PImpl) ReadFromDB() {
	fmt.Println("vim-go")
}
