package publisher

import "testing"

func TestPublishOk(t *testing.T) {
	p := &PImpl{}
	p.Publish()
}

func TestReadFromDB(t *testing.T) {
	p := &PImpl{}
	p.ReadFromDB()
}
