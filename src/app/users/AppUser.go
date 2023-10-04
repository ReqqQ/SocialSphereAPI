package appuser

import infrastructurebus "github.com/ReqqQ/SocialSphereAPI/src/infrastructure/bus"

type AppUserInterface interface {
	GetUser(id int)
}
type AppUser struct {
	QueryHandler infrastructurebus.QueryHandler
}

type GetUserQuery struct {
	UserId int
}

func (q GetUserQuery) QueryName() string {
	return "GetUserQuery"
}

func (a AppUser) GetUser(id int) {
	query := GetUserQuery{UserId: id}

	a.QueryHandler.QueryHandle(query)
}
