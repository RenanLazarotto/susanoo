package emprego

import (
	"github.com/gofiber/fiber/v2"

	empregoHandler "tsukuyomi/handlers/emprego"
	"tsukuyomi/repositories"
	empregoRepository "tsukuyomi/repositories/emprego"
	empregoService "tsukuyomi/services/emprego"
)

func RegisterRoutes(app *fiber.App, repository repositories.Repository) {
	empregoRepository := empregoRepository.NewRepository(repository)
	empregoService := empregoService.NewService(empregoRepository)

	handler := empregoHandler.NewHandler(empregoService)

	router := app.Group("/emprego")
	router.Post("/", handler.Create)
	router.Get("/", handler.FindAll)
	router.Get("/:id", handler.FindByID)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)
}
