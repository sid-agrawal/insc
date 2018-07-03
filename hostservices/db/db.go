package db

import "fmt"

type Db struct {
	something interface{}
}

func (d *Db) Read() {
	fmt.Println("vim-go")
}

func (d *Db) Write() {
	fmt.Println("vim-go")
}
