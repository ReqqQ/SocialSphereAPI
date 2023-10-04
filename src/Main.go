package main

import (
	"github.com/ReqqQ/SocialSphereAPI/config"
)

func main() {
	app := config.Init()
	config.InitCQRS()
	//config.App.GetCommands().GetSyncDBCommand().SyncDB()
	config.App.GetControllers().InitGetRoutes(app)
	config.Start(app)
}
