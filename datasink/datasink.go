// Package datasink presents the interface that every data consumer has to implement
package datasink

type datasink interface {
	Connect(string) error
	SubscribeTopic(string) error
	Close()
}
