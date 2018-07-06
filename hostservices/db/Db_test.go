package db

import "testing"

func TestReaderOk(t *testing.T) {
	d := &BoltDB{}
	d.Read()
}

func TestWriteOk(t *testing.T) {
	d := &BoltDB{}
	d.Write()
}
