package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"tsukuyomi/config"
	_ "tsukuyomi/docs"
	"tsukuyomi/repositories"
	contatoEmpresa "tsukuyomi/routers/contato_empresa"
	"tsukuyomi/routers/empresa"
	"tsukuyomi/routers/endereco"
	enderecoEmpresa "tsukuyomi/routers/endereco_empresa"
)

func SetupRouter(app *fiber.App, config *config.Config) {
	// Rota de documentação
	app.Get("/swagger/*", swagger.HandlerDefault)

	repository := repositories.NewRepository(config)

	empresa.RegisterRoutes(app, repository)
	endereco.RegisterRoutes(app, repository)
	contatoEmpresa.RegisterRoutes(app, repository)
	enderecoEmpresa.RegisterRoutes(app, repository)
}
