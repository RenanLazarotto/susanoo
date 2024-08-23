package empresa

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"tsukuyomi/models"
	empresaService "tsukuyomi/services/empresa"
)

type EmpresaHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}
type empresaHandler struct {
	Service empresaService.Service
}

func NewHandler(service empresaService.Service) EmpresaHandler {
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

func (h *empresaHandler) GetAll(c *fiber.Ctx) error {
	criteria := map[string]interface{}{}
	c.BodyParser(&criteria)

	result, err := h.Service.GetAll(criteria)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Meta:  models.Meta{Count: 0},
			Error: err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Meta: models.Meta{Count: len(result)},
		Data: result,
	})
}

func (h *empresaHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.Service.GetByID(id)
	if err != nil {
		if ok := errors.Is(err, gorm.ErrRecordNotFound); ok {
			return c.Status(fiber.StatusNotFound).JSON(models.Response{
				Meta: models.Meta{Count: 0},
			})
		}
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

func (h *empresaHandler) Update(c *fiber.Ctx) error {
	empresa := new(models.Empresa)
	c.BodyParser(&empresa)

	if empresa.ID == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Meta:  models.Meta{Count: 0},
			Error: "id is required",
		})
	}

	result, err := h.Service.Update(*empresa)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Meta:  models.Meta{Count: 0},
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Meta: models.Meta{Count: 1},
		Data: result,
	})
}
