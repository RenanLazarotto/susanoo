package endereco

import (
	"github.com/gofiber/fiber/v2"

	enderecoHandler "tsukuyomi/handlers/endereco"
	"tsukuyomi/repositories"
	enderecoRepository "tsukuyomi/repositories/endereco"
	enderecoService "tsukuyomi/services/endereco"
)

func RegisterRoutes(app *fiber.App, repository repositories.Repository) {
	enderecoRepository := enderecoRepository.NewRepository(repository)
	enderecoService := enderecoService.NewService(enderecoRepository)

	handler := enderecoHandler.NewHandler(enderecoService)

	router := app.Group("/endereco")
	router.Post("/", handler.Create)
	router.Get("/", handler.FindAll)
	router.Get("/:id", handler.FindByID)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)

}
