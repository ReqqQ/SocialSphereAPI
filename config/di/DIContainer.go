//go:generate wire
package di

import (
	"github.com/google/wire"

	"github.com/ReqqQ/SocialSphereAPI/src/ui"
)

type AppService struct {
	Controllers *ui.Controllers
}

var (
	AppServiceSet = wire.NewSet(provideControllers, provideAppService)
)

func InitDIContainer() (*AppService, error) {
	wire.Build(AppServiceSet)
	return &AppService{}, nil
}

func provideControllers() *ui.Controllers {
	return &ui.Controllers{}
}

func provideAppService(controllers *ui.Controllers) *AppService {
	return &AppService{
		Controllers: controllers,
	}
}
