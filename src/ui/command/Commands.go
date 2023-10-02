package uicommand

type CommandsInterface interface {
	GetSyncDBCommand() SyncDbInterface
}
type Commands struct {
	syncDbCommand SyncDbCommand
}

func (c Commands) GetSyncDBCommand() SyncDbInterface {
	return c.syncDbCommand
}
