package db

import "testing"

func TestReaderOk(t *testing.T) {
	d := &Db{}
	d.Read()
}

func TestiWriteOk(t *testing.T) {
	d := &Db{}
	d.Write()
}
