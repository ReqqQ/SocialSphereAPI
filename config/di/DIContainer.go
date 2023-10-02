//go:generate wire
package di

import (
	"github.com/google/wire"

	uicommands "github.com/ReqqQ/SocialSphereAPI/src/ui/command"
	uicontrollers "github.com/ReqqQ/SocialSphereAPI/src/ui/controllers"
)

type AppServiceInterface interface {
	GetCommands() uicommands.CommandsInterface
	GetControllers() uicontrollers.ControllersInterface
}

type AppService struct {
	Controllers uicontrollers.Controllers
	Commands    uicommands.Commands
}

var (
	AppServiceSet = wire.NewSet(provideControllers, provideAppService, provideCommands)
)

func InitDIContainer() (*AppService, error) {
	wire.Build(AppServiceSet)
	return &AppService{}, nil
}

func (a AppService) GetCommands() uicommands.CommandsInterface {
	return a.Commands
}

func (a AppService) GetControllers() uicontrollers.ControllersInterface {
	return a.Controllers
}

func provideControllers() uicontrollers.Controllers {
	return uicontrollers.Controllers{}
}
func provideCommands() uicommands.Commands {
	return uicommands.Commands{}
}
func provideAppService(controllers uicontrollers.Controllers, commands uicommands.Commands) *AppService {
	return &AppService{
		Controllers: controllers,
		Commands:    commands,
	}
}
