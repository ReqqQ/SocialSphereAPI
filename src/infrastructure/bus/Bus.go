package infrastructurebus

var QueryHandlerList map[string]RepositoryHandlerInterface

type QueryHandlerInterface interface {
	QueryHandle(q QueryInterface) interface{}
}

type RepositoryHandlerInterface interface {
	Invoke() interface{}
}

type CommandHandlerInterface interface {
}
type QueryInterface interface {
	QueryName() string
}

type QueryHandler struct {
}

type CommandHandler struct{}

func (qh QueryHandler) QueryHandle(q QueryInterface) interface{} {
	repository := QueryHandlerList[q.QueryName()]

	return repository.Invoke()
}
