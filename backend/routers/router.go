package routers

import (
	"github.com/gofiber/fiber/v2"

	"tsukuyomi/config"
	"tsukuyomi/repositories"
	"tsukuyomi/routers/empresa"
)

func SetupRouter(app *fiber.App, config *config.Config) {
	repository := repositories.NewRepository(config)

	empresa.RegisterRoutes(app, repository)
}
