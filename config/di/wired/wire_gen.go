// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wired

import (
	"github.com/google/wire"

	"github.com/ReqqQ/SocialSphereAPI/src/ui"
)

// Injectors from DIContainer.go:

func InitDIContainer() (*AppService, error) {
	controllers := provideControllers()
	appService := provideAppService(controllers)
	return appService, nil
}

// DIContainer.go:

type AppService struct {
	Controllers *ui.Controllers
}

var (
	AppServiceSet = wire.NewSet(provideControllers, provideAppService)
)

func provideControllers() *ui.Controllers {
	return &ui.Controllers{}
}

func provideAppService(controllers *ui.Controllers) *AppService {
	return &AppService{
		Controllers: controllers,
	}
}
