package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/ReqqQ/SocialSphereAPI/config/di/wired"
)

var (
	App *wired.AppService
	err error
)

func Init() *fiber.App {
	App, err = wired.InitDIContainer()
	if err != nil {
		return nil
	}
	err = godotenv.Load(".env")

	return fiber.New()
}

func Start(app *fiber.App) {
	app.Listen(":3000")
}
