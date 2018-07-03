package db

import "fmt"

// DB is used to connect to boltdb
type DB struct {
	something interface{}
}

func (d *DB) Read() {
	fmt.Println("vim-go")
}

func (d *DB) Write() {
	fmt.Println("vim-go")
}
