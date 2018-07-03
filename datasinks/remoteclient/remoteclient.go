package remoteclient

import "fmt"

// RemoteClient implements the datasinks interface.
type RemoteClient struct {
	something interface{}
}

// Connect is used to connect to the Publisher of data.
func (r *RemoteClient) Connect() {
	fmt.Println("vim-go")
}
