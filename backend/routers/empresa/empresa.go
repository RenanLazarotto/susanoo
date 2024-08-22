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
	router.Post("/create", handler.Create)
}
