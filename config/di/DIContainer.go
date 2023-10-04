//go:generate wire
package di

import (
	"github.com/google/wire"

	infrastructurebus "github.com/ReqqQ/SocialSphereAPI/src/infrastructure/bus"
	uicommands "github.com/ReqqQ/SocialSphereAPI/src/ui/command"
	uicontrollers "github.com/ReqqQ/SocialSphereAPI/src/ui/controllers"
)

type AppServiceInterface interface {
	GetCommands() uicommands.CommandsInterface
	GetControllers() uicontrollers.ControllersInterface
	GetQueryHandlers() infrastructurebus.QueryHandlerInterface
	GetCommandHandlers() infrastructurebus.CommandHandlerInterface
}

type AppService struct {
	Controllers    uicontrollers.Controllers
	Commands       uicommands.Commands
	QueryHandler   infrastructurebus.QueryHandler
	CommandHandler infrastructurebus.CommandHandler
}

var (
	AppServiceSet = wire.NewSet(provideControllers, provideAppService, provideCommands, provideQueryHandler, provideCommandHandler)
)

func InitDIContainer() (*AppService, error) {
	wire.Build(AppServiceSet)
	return &AppService{}, nil
}

func (a AppService) GetCommands() uicommands.CommandsInterface {
	return a.Commands
}

func (a AppService) GetQueryHandlers() infrastructurebus.QueryHandlerInterface {
	return a.QueryHandler
}

func (a AppService) GetCommandHandlers() infrastructurebus.CommandHandlerInterface {
	return a.CommandHandler
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

func provideQueryHandler() infrastructurebus.QueryHandler {
	return infrastructurebus.QueryHandler{}
}
func provideCommandHandler() infrastructurebus.CommandHandler {
	return infrastructurebus.CommandHandler{}
}

func provideAppService(
	controllers uicontrollers.Controllers,
	commands uicommands.Commands,
	queryHandler infrastructurebus.QueryHandler,
	commandHandler infrastructurebus.CommandHandler,
) *AppService {
	return &AppService{
		Controllers:    controllers,
		Commands:       commands,
		QueryHandler:   queryHandler,
		CommandHandler: commandHandler,
	}
}
