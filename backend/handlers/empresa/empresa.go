package empresa

import (
	"github.com/gofiber/fiber/v2"

	"tsukuyomi/models"
	es "tsukuyomi/services/empresa"
)

type EmpresaHandler interface {
	Create(c *fiber.Ctx) error
}
type empresaHandler struct {
	Service es.Service
}

func NewHandler(service es.Service) EmpresaHandler {
	return &empresaHandler{
		Service: service,
	}
}

func (h *empresaHandler) Create(c *fiber.Ctx) error {
	empresa := new(models.Empresa)

	c.BodyParser(&empresa)

	result, err := h.Service.Create(*empresa)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Meta:  models.Meta{Count: 0},
			Error: err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Meta: models.Meta{Count: 1},
		Data: result,
	})
}
