// Package datasource presents the interface that every data producer has to implement.
package datasource

type datasource interface {
	Connect(string) error
	Close()
	// There is no Push since data will always be pulled from the source.
}
