package db

import "fmt"

// DB is the API for the package
type DB interface {
	Read()
	Write()
}

// BoltDB is used to connect to boltdb
type BoltDB struct {
	something interface{}
}

func (d *BoltDB) Read() {
	fmt.Println("vim-go")
}

func (d *BoltDB) Write() {
	fmt.Println("vim-go")
}
