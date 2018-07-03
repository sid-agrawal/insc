package db

import "testing"

func TestReaderOk(t *testing.T) {
	d := &DB{}
	d.Read()
}

func TestWriteOk(t *testing.T) {
	d := &DB{}
	d.Write()
}
