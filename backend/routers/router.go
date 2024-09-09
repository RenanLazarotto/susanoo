package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"tsukuyomi/config"
	_ "tsukuyomi/docs"
	"tsukuyomi/repositories"
	"tsukuyomi/routers/empresa"
	"tsukuyomi/routers/endereco"
)

func SetupRouter(app *fiber.App, config *config.Config) {
	// Rota de documentação
	app.Get("/swagger/*", swagger.HandlerDefault)

	repository := repositories.NewRepository(config)

	empresa.RegisterRoutes(app, repository)
	endereco.RegisterRoutes(app, repository)
}
