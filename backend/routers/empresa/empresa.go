package empresa

import (
	"github.com/gofiber/fiber/v2"

	empresaHandler "tsukuyomi/handlers/empresa"
	"tsukuyomi/repositories"
	empresaRepository "tsukuyomi/repositories/empresa"
	empresaService "tsukuyomi/services/empresa"
)

func RegisterRoutes(app *fiber.App, repository repositories.Repository) {
	empresaRepository := empresaRepository.NewRepository(repository)
	empresaService := empresaService.NewService(empresaRepository)

	handler := empresaHandler.NewHandler(empresaService)

	router := app.Group("/empresa")
	router.Post("/", handler.Create)
	router.Get("/", handler.GetAll)
	router.Get("/:id", handler.GetByID)
	router.Put("/", handler.Update)
}
