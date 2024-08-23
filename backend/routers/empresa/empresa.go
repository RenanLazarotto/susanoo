package empresa

import (
	"github.com/gofiber/fiber/v2"

	eh "tsukuyomi/handlers/empresa"
	"tsukuyomi/repositories"
	er "tsukuyomi/repositories/empresa"
	es "tsukuyomi/services/empresa"
)

func RegisterRoutes(app *fiber.App, repository repositories.Repository) {
	empresaRepository := er.NewRepository(repository)
	empresaService := es.NewService(empresaRepository)

	handler := eh.NewHandler(empresaService)

	router := app.Group("/empresa")
	router.Post("/", handler.Create)
	router.Get("/", handler.GetAll)
	router.Get("/:id", handler.GetByID)
	router.Put("/", handler.Update)
}
