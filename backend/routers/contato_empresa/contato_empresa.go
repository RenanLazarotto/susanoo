package contatoempresa

import (
	"github.com/gofiber/fiber/v2"

	contatoEmpresaHandler "tsukuyomi/handlers/contato_empresa"
	"tsukuyomi/repositories"
	contatoEmpresaRepository "tsukuyomi/repositories/contato_empresa"
	contatoEmpresaService "tsukuyomi/services/contato_empresa"
)

func RegisterRoutes(app *fiber.App, repository repositories.Repository) {
	contatoEmpresaRepository := contatoEmpresaRepository.NewRepository(repository)
	contatoEmpresaService := contatoEmpresaService.NewService(contatoEmpresaRepository)

	handler := contatoEmpresaHandler.NewHandler(contatoEmpresaService)

	router := app.Group("/contato-empresa")
	router.Post("/", handler.Create)
	router.Get("/", handler.FindAll)
	router.Get("/:id", handler.FindByID)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)

}
