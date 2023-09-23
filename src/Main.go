package main

import (
	"github.com/ReqqQ/SocialSphereAPI/config"
)

func main() {
	app := config.Init()
	config.App.Controllers.InitGetRoutes(app)
	config.Start(app)
}
