package endereco

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"tsukuyomi/models"
	enderecoService "tsukuyomi/services/endereco"
)

type EnderecoHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}
type enderecoHandler struct {
	Service enderecoService.Service
}

func NewHandler(service enderecoService.Service) EnderecoHandler {
	return &enderecoHandler{
		Service: service,
	}
}

func (h *enderecoHandler) Create(c *fiber.Ctx) error {
	endereco := new(models.Endereco)

	c.BodyParser(&endereco)

	result, err := h.Service.Create(*endereco)
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

func (h *enderecoHandler) GetAll(c *fiber.Ctx) error {
	criteria := map[string]interface{}{}
	c.BodyParser(&criteria)

	result, err := h.Service.GetAll(criteria)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Meta:  models.Meta{Count: 0},
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Meta: models.Meta{Count: len(result)},
		Data: result,
	})
}

func (h *enderecoHandler) GetByID(c *fiber.Ctx) error {
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
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Meta: models.Meta{Count: 1},
		Data: result,
	})
}

func (h *enderecoHandler) Update(c *fiber.Ctx) error {
	endereco := new(models.Endereco)
	c.BodyParser(&endereco)

	if endereco.ID == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Meta:  models.Meta{Count: 0},
			Error: "id is required",
		})
	}

	result, err := h.Service.Update(*endereco)
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
