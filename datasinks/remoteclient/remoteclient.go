package remoteclient

import "fmt"

type RemoteClient struct {
	something interface{}
}

func (r *RemoteClient) Connect() {
	fmt.Println("vim-go")
}
