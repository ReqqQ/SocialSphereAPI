package uicommand

import "fmt"

type SyncDbInterface interface {
	SyncDB()
}
type SyncDbCommand struct {
}

func (s SyncDbCommand) SyncDB() {
	fmt.Println("echo")
	//ctx := context.Background()
	for {
		//go fetchUsers()
	}
}
