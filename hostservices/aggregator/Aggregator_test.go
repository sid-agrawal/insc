package aggregator

import "testing"

func TestRun(t *testing.T) {
	a := &AImpl{}
	a.Run()
}

func TestPull(t *testing.T) {
	a := &AImpl{}
	a.Pull()
}

func TestAddSource(t *testing.T) {
	a := &AImpl{}
	a.AddSource()
}

func TestWriteToDB(t *testing.T) {
	a := &AImpl{}
	a.WriteToDB()
}
