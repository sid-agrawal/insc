package aggregator

import "fmt"

// Aggregator is the API for the package
type Aggregator interface {
	Run()
	AddSource()
	Pull()
	Write()
}

// AImpl implements ...
type AImpl struct {
	something interface{}
	// DB Handle
}

// Run the Aggregator
func (a *AImpl) Run() {
	fmt.Println("vim-go")
}

// Pull Data from different sources
func (a *AImpl) Pull() {
	fmt.Println("vim-go")
}

// AddSource will add a data source. For example, an IP or IPC
func (a *AImpl) AddSource() {
	fmt.Println("vim-go")
}

// WriteToDB will write all the data to the DB.
// It will call the Write method of DB
func (a *AImpl) WriteToDB() {
	fmt.Println("vim-go")
}
